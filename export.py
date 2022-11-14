import argparse
import re
import shutil
import sys
from pathlib import Path

import frontmatter
from loguru import logger
from obsidian_to_hugo.wiki_links_processor import replace_wiki_links


class ObsidianPage:
    def __init__(self, filepath: Path, post: frontmatter.Post):
        self.filepath = filepath
        self.post = post

    def process(self):
        self.gist()
        self.post.content = replace_wiki_links(self.post.content)

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

    return parser.parse_args()


def gather(obsidian_vault: Path) -> list[ObsidianPage]:
    filepaths = obsidian_vault.glob("**/*.md")
    obsidian_pages = []

    for filepath in filepaths:
        with open(filepath, "r") as f:
            post = frontmatter.loads(f.read())
            if post.metadata.get("draft", False):
                continue
            obsidian_pages.append(ObsidianPage(filepath=filepath, post=post))
    return obsidian_pages


def rm_tree(pth: Path):
    if not pth.exists():
        return
    logger.info(f"Cleaning up (removing) {pth}")
    pth = Path(pth)
    for child in pth.glob("*"):
        if child.is_file():
            child.unlink()
        else:
            rm_tree(child)
    pth.rmdir()


def write(src: Path, dst: Path, obsidian_pages: list[ObsidianPage]):
    rm_tree(dst)

    for obsidian_page in obsidian_pages:
        target = Path(
            str(obsidian_page.filepath.absolute()).replace(str(src), str(dst))
        )
        target.parent.mkdir(parents=True, exist_ok=True)
        with open(target, "w") as f:
            f.write(frontmatter.dumps(obsidian_page.post))
            logger.info(f"Wrote {target}")


def copy_static_files(src: Path, dst: Path):
    rm_tree(dst)
    logger.info(f"Copying static files into {dst}...")
    shutil.copytree(src=src, dst=dst)
    logger.info("Done copying static files.")


def main():
    args = parse_args()
    obsidian_pages = gather(args.obsidian_vault)
    for obsidian_page in obsidian_pages:
        obsidian_page.process()
    write(
        src=args.obsidian_vault, dst=args.hugo_contents, obsidian_pages=obsidian_pages
    )
    copy_static_files(
        src=args.obsidian_vault / "static/", dst=args.hugo_contents.parent / "static"
    )
    logger.info("Done.")


if __name__ == "__main__":
    sys.exit(main())
