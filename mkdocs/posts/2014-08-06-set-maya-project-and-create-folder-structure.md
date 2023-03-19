---
date: 2014-08-06
tags:
- maya
---

# Set Maya project and create folder structure

After having done a quick Google search, it seems nobody has yet posted a quick Python snippet that does this.

<!-- more -->

... so without further ado:

```python
import maya.cmds as cmds
import maya.mel as mel

def create_folder( directory ):
    if not os.path.exists( directory ):
        os.makedirs( directory )

maya_dir = '//server/share/path/to/maya'
create_folder( maya_dir )

mel.eval('setProject \"' + maya_dir + '\"')

for file_rule in cmds.workspace(query=True, fileRuleList=True):
    file_rule_dir = cmds.workspace(fileRuleEntry=file_rule)
    maya_file_rule_dir = os.path.join( maya_dir, file_rule_dir)
    create_folder( maya_file_rule_dir )
```