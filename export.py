import argparse
import re
import shutil
import sys
from pathlib import Path

import frontmatter
from loguru import logger
from obsidian_to_hugo.wiki_links_processor import replace_wiki_links
from yaml.scanner import ScannerError


class ObsidianPage:
    def __init__(self, filepath: Path, post: frontmatter.Post):
        self.filepath = filepath
        self.post = post

    def process(self, args: argparse.Namespace):
        self.gist()
        self.post.content = self.replace_static_vault_links(
            self.post.content,
            src_uri=args.obsidian_vault_static_uri,
            dst_uri=args.hugo_static_uri,
        )
        self.post.content = replace_wiki_links(self.post.content)
        self.jupyter()

    def replace_static_vault_links(
        self, content: str, src_uri: str, dst_uri: str
    ) -> str:
        """Replace Obsidian vault static links to Hugo static links."""

        markdown_link = r"\[(.*)\]\((.*)\)"
        markdown_image = r"\!\[(.*)\]\((.*)\)"
        patterns = [markdown_link, markdown_image]

        for pattern in patterns:
            match = re.finditer(pattern, self.post.content)
            if match:
                for group in match:
                    link_text = group[1]
                    link = group[2]

                    updated_link = link.replace(src_uri, dst_uri)

                    if pattern == markdown_link:
                        content = content.replace(
                            f"[{link_text}]({link})", f"[{link_text}]({updated_link})"
                        )
                    elif pattern == markdown_image:
                        content = content.replace(
                            f"![{link_text}]({link})", f"[{link_text}]({updated_link})"
                        )
        return content

    def gist(self):
        """Convert gist into a shortcode.

        Notes:
            - Obsidian requirement: https://github.com/linjunpop/obsidian-gist
              Supports gists with or without a filename, but not without username.
            - Shortcode: https://gohugo.io/content-management/shortcodes/#gist
        """
        pattern = r"(```gist\s((.*)/(.*))\s```)"
        match = re.finditer(pattern, self.post.content)
        if match:
            for group in match:
                code_block = group[1]
                # gist_data = group[2]
                username = group[3]
                gist_id = group[4]

                if "#" in gist_id:
                    gist_id, filename = gist_id.split("#")
                    filename = f'"{filename}"'

                    shortcode = f"{{{{< gist {username} {gist_id} {filename} >}}}}"
                else:
                    shortcode = f"{{{{< gist {username} {gist_id} >}}}}"

                self.post.content = self.post.content.replace(code_block, shortcode)
                logger.info(f"Replaced gist in {self.filepath}")

    def jupyter(self):
        self.post.content = self.post.content.replace("```jupyter", "```python")


def parse_args():
    parser = argparse.ArgumentParser(
        formatter_class=argparse.ArgumentDefaultsHelpFormatter
    )
    parser.add_argument(
        "--obsidian-vault",
        type=Path,
        default=Path("obsidian"),
        help="Obsidian source folder",
    )
    parser.add_argument(
        "--hugo-contents",
        type=Path,
        default=Path("hugo/content"),
        help="Hugo contents folder",
    )
    parser.add_argument(
        "--log-level",
        type=str,
        default="INFO",
        help="Log level. One of: DEBUG, INFO, WARNING, ERROR, CRITICAL",
    )
    parser.add_argument(
        "--copy-static-files",
        type=bool,
        default=True,
        help="Copy static files from Obsidian vault to Hugo static folder",
    )
    parser.add_argument(
        "--obsidian-vault-static-uri",
        type=str,
        default="fredrikaverpil.github.io/obsidian/static",
        help="Obsidian vault's static folder URI",
    )
    parser.add_argument(
        "--hugo-static-uri",
        type=str,
        default="/static",
        help="Hugo's static folder URI",
    )

    return parser.parse_args()


def set_log_level(level: str) -> None:
    logger.remove()
    logger.add(sink=sys.stderr, level=level)


def gather(obsidian_vault: Path) -> list[ObsidianPage]:
    filepaths = obsidian_vault.glob("**/*.md")
    obsidian_pages = []

    for filepath in filepaths:
        with open(filepath, "r") as f:
            try:
                post = frontmatter.loads(f.read())
            except ScannerError as err:
                logger.error(f"Error parsing frontmatter in {filepath}: {err}")
                continue
            if post.metadata.get("draft", False):
                logger.warning(f"Skipping draft: {filepath}")
                continue
            obsidian_pages.append(ObsidianPage(filepath=filepath, post=post))
    return obsidian_pages


def rm_tree(pth: Path):
    if not pth.exists():
        return
    logger.debug(f"Cleaning up (removing) {pth}")
    pth = Path(pth)
    for child in pth.glob("*"):
        if child.is_file():
            child.unlink()
        else:
            rm_tree(child)
    pth.rmdir()


def write(src: Path, dst: Path, obsidian_pages: list[ObsidianPage]):
    logger.info(f"Copying {len(obsidian_pages)} markdown files from {src} to {dst}...")
    rm_tree(dst)
    file_counter = 0
    for obsidian_page in obsidian_pages:
        target = dst / obsidian_page.filepath.relative_to(src)

        target.parent.mkdir(parents=True, exist_ok=True)

        with open(target, "w") as f:
            f.write(frontmatter.dumps(obsidian_page.post))
            file_counter += 1

    logger.info(f"Done copying {file_counter} markdown files.")


def copy_static_files(src: Path, dst: Path):
    rm_tree(dst)
    logger.info(f"Copying static files into {dst}...")
    shutil.copytree(src=src, dst=dst)
    logger.info("Done copying static files.")


def main():
    args = parse_args()
    set_log_level(level=args.log_level)
    obsidian_pages = gather(args.obsidian_vault)
    for obsidian_page in obsidian_pages:
        obsidian_page.process(args)
    write(
        src=args.obsidian_vault, dst=args.hugo_contents, obsidian_pages=obsidian_pages
    )
    if args.copy_static_files:
        copy_static_files(
            src=args.obsidian_vault / "static/",
            dst=args.hugo_contents.parent / "static",
        )

    logger.info("Done.")


if __name__ == "__main__":
    sys.exit(main())
