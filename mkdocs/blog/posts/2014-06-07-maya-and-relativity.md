---
date: 2014-06-07
tags:
- maya
---

# Maya and relativity

A short note on how to work with relative paths in Autodesk’s Maya.

<!-- more -->

### Sourceimages

You can define a relative path (from where your Maya project is set) using workspace.mel or the “File -> Project Window”. Here, set the sourceimages folder to e.g. `../textures`. It is important that you do not use the word “sourceimages” in the path or this won’t work.

For any file texture that you wish loaded from there, enter sourceimages/myTex.tif in the filepath of the file node.

If you save and open the Maya ASCII scene file, this relative filepath will indeed be the filepath written into the file. This means if you want to move the sourceimages folder, all you have to edit is the workspace.mel (or editing via the Project Window).

### References

After having created a reference, you need to edit the “Unresolved name” of the reference into e.g. `../references/myReference.ma` in order to set it relatively (from where your Maya project is set).

If you save and open the Maya ASCII scene file, this relative filepath will indeed be the filepath written into the file.

### V-Ray nodes

V-Ray nodes such as for proxies supports relative paths out of the box. Just set the path to e.g. your vrmesh to `../proxies/myProxy.vrmesh`. This is all relative to where your Maya project is set.

### Environment variables

Maya can also take environment variables and use inside of a file path. Just set the “Unresolved name” of a reference to e.g. `../../$MY_ENV/maya/scenes/myReference.ma` to have it load the reference from this location. If the environment variable $MY_ENV is not found, Maya will look into its own project (-proj flag) for the file.