---
date: 2012-08-18
authors:
  - fredrikaverpil
comments: true
tags:
- nuke
- python
---

# Metadata and EXIF in Nuke

In [The Foundry’s Nuke](http://www.thefoundry.co.uk/products/nuke/), accessing the metadata of an image sequence’s Read node can be done via the ViewMetaData node. But if you wish to extract values out of the metadata and e.g. burn the timestamp for each frame into a render, it is actually easier to just do some python scripting.

<!-- more -->

In the example below, I am using JPEG footage from a GoPro HD Hero2 but this technique also works fine for Cineon and DPX files etc.

## Find the metadata to extract

First, check the name of the metadata you wish to extract. This can be done via the ViewMetaData node.

![](/static/nuke_metadata/metadata_viewmetadatanode.png)
*In this case, I am looking for the image’s date and time values.*

## Extracting the metadata using Python

Let’s say we wish to extract the value of the metadata key called exif/0/DateTime (which, in this case, is the date and time of when the image file was recorded) and print this timestamp onto each frame throughout the image sequence.

In my case I am going to attach a Text node directly to the Read node and type the following into the Text node’s “message” field:

```python
Current time: [python {nuke.thisNode().metadata()['exif/0/DateTime']}]
```

![](/static/nuke_metadata/metadata_textnode.png)
*This would cause the resulting render to say “Current time: 2012:08:10 15:51:13” for this frame.*

## Defining a Read node to extract metadata from

If you wish to access metadata from a Read node outside of your node tree, you can instead define exactly which Read node you are reading metadata from, using nuke.toNode rather than nuke.thisNode. In the example below I am accessing the value of the metadata key exif/0/DateTime of a Read node called “Read1”:

```python
Current time: [python {nuke.toNode( "Read1" ).metadata()['exif/0/DateTime']}]
```