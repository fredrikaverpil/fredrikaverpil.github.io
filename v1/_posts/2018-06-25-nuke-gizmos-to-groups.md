---
layout: post
title: Nuke gizmos to groups
tags: [nuke, python]
---

Recursively replace all gizmos in Nuke script with a group. The only exception is the `Cryptomatte` nodes, which will be maintained as gizmos.

<!--more-->

Copy-paste the below code into the Nuke script editor and run.


```python

import uuid
import nuke


def is_gizmo(node):
    """Return True if given node is a gizmo (and not allowed gizmo type)"""

    allowed_gizmo_classes = ('Cryptomatte')

    for knob in node.knobs():
        if 'gizmo' in knob and node.Class() not in allowed_gizmo_classes:
            return True


def get_gizmo_names():
    """Return the fullName attribute for all gizmos in script"""
    
    gizmos = []
    all_nodes = nuke.allNodes(recurseGroups=True)

    if all_nodes:
        gizmos = [node.fullName() for node in all_nodes if is_gizmo(node)]

    return gizmos


def deselect_all_nodes():
    """De-select all nodes"""

    for i in nuke.allNodes(recurseGroups=True):
        i.knob('selected').setValue(False)


def convert_gizmo_to_group(gizmo_full_name):
    """Convert given gizmo (gizmo.fullName) to group"""
    
    gizmo = nuke.toNode(gizmo_full_name)

    inputs = []
    for x in range(0, gizmo.maximumInputs()):
        if gizmo.input(x):
            inputs.append(gizmo.input(x))
        else:
            inputs.append(False)

    original_name = gizmo.knob('name').value()
    xpos = gizmo.xpos()
    ypos = gizmo.ypos()
    uid_name = uuid.uuid4()

    gizmo.knob('name').setValue('%s' % uid_name)
    deselect_all_nodes()
    gizmo.knob('selected').setValue(True)

    with gizmo:
        new_group = gizmo.makeGroup()

        deselect_all_nodes()
        nuke.delete(gizmo)

        new_group.knob('name').setValue(original_name)
        new_group['xpos'].setValue(xpos)
        new_group['ypos'].setValue(ypos)

        for x in range(0, new_group.maximumInputs()):
            new_group.setInput(x, None)
            if inputs[x]:
                new_group.connectInput(x, inputs[x])


def main():
    """Main script"""
    # Store the current selection
    current_selection = nuke.selectedNodes()

    while get_gizmo_names():
        gizmo_full_name = get_gizmo_names()[0]
        convert_gizmo_to_group(gizmo_full_name)
        print('Converted %s' % gizmo_full_name)

    # Restore original selection
    for n in current_selection:
        try:
            n['selected'].setValue(True)
        except:
            pass

main()
```
