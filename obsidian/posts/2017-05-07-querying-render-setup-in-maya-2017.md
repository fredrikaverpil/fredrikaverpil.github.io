---
title: Querying Render Setup in Maya 2017
tags: [python, maya]
draft: false

# PaperMod
ShowToc: false
TocOpen: false

created: 2017-05-07T02:00:12+02:00
updated: 2022-11-15T22:29:17+01:00
---

It seems Autodesk did not create a Render Setup documentation. This is me collecting code snippets and exploring Python functions.




## MEL code snippets

```
# Return render layer names without the "rs_" prefix
$renderLayerNames = `renderSetup -q -renderLayers`;
```


## Python code snippets


### Render layer names

```python
import maya.cmds as cmds

# Return render layer names without the "rs_" prefix
render_layer_names = cmds.renderSetup(q=True, renderLayers=True)
```

### Render layer names and their renderability

```python
import maya.app.renderSetup.model.renderSetup as renderSetup

render_setup = renderSetup.instance()
render_layers = render_setup.getRenderLayers()

for render_layer in render_layers:
    render_layer_name = render_layer.name()  # Without "rs_" prefix
    is_renderable = render_layer.isRenderable()
    print(render_layer_name, is_renderable)
```


## Exploring Render Setup

If you run the following, you're going to get all the callable attributes of `render_setup` (such as strings, methods etc):

```python
import maya.app.renderSetup.model.renderSetup as renderSetup

def get_callable_attributes(obj):
    """Returns the callable attributes of an object"""
    
    callable_attrs = {}
    for attr in dir(obj):
        attr_obj = getattr(obj, attr)
        if callable(attr_obj):
            callable_attrs[attr] = attr_obj
    return sorted(callable_attrs)

render_setup = renderSetup.instance()
callable_attrs = get_callable_attributes(render_setup)
for callable_attr in callable_attrs:
    print(callable_attr)
```

This is what I get with Maya 2017 Update 3:

```
__class__
__delattr__
__format__
__getattribute__
__hash__
__init__
__new__
__reduce__
__reduce_ex__
__repr__
__setattr__
__sizeof__
__str__
__subclasshook__
__weakref__
_afterDuplicate
_afterLoadReferenceCB
_afterOpenCB
_afterUnloadReferenceCB
_beforeDuplicate
_beforeLoadReferenceCB
_beforeSaveSceneCB
_beforeUnloadReferenceCB
_cleanObservers
_decodeChildren
_decodeProperties
_encodeProperties
_getBackAttr
_getFrontAttr
_getListItemsAttr
_getNotesPlug
_hasNotesPlug
_notifyActiveLayerObservers
_onMayaNodeAddedCB
_onNodeRemoved
_preRenderLayerDelete
_switchToLayerFileIO
acceptImport
addActiveLayerObserver
addAttribute
addExternalContentForFileAttr
addListObserver
ancestors
appendChild
appendRenderLayer
attachChild
attachRenderLayer
attributeAffects
clearAll
clearListObservers
compute
connectionBroken
connectionMade
copyInternalData
createRenderLayer
creator
decode
dependsOn
detachChild
detachRenderLayer
dispose
doNotWrite
encode
forceCache
getBack
getChildren
getDefaultRenderLayer
getExternalContent
getFilesToArchive
getFront
getInternalValueInContext
getNotes
getRenderLayer
getRenderLayers
getVisibleRenderLayer
hasActiveLayerObserver
inheritAttributesFrom
initListItems
initializer
internalArrayCount
isAbstractClass
isAcceptableChild
isPassiveOutput
itemAdded
itemRemoved
legalConnection
legalDisconnection
name
parent
passThroughToMany
passThroughToOne
postConstructor
removeActiveLayerObserver
removeListObserver
setBack
setDependentsDirty
setDoNotWrite
setExternalContent
setExternalContentForFileAttr
setFront
setInternalValueInContext
setMPSafe
setNotes
shouldSave
switchToLayer
switchToLayerUsingLegacyName
thisMObject
type
typeId
typeName
```

By accessing the built-in [`help()`](https://docs.python.org/2/library/functions.html#help) Python function could shed some additional light... This (among other things) will print the function's docstring (`function.__doc__`) if available.

The following prints all "callables" and any documentation in markdown which can be copy-pasted to a gist for (some) readability. Please note that you're most likely just interested in objects of type `instancemethod`. I'm deliberately skipping any attributes starting with an underscore.

```python
import maya.app.renderSetup.model.renderSetup as renderSetup


def get_callable_attributes(obj):
    """Returns the callable attributes of an object"""
    
    callable_attrs = {}
    for attr in dir(obj):
        attr_obj = getattr(obj, attr)
        if callable(attr_obj):
            callable_attrs[attr] = attr_obj
    return callable_attrs


def print_markdown(callable_attrs):
    """Print the help of each callable attribute in markdown"""

    # Index
    print('## Index')
    for attr_name in sorted(callable_attrs):
            if not attr_name.startswith('_'):
                print('<a href="#' + attr_name + '">`' + attr_name + '`</a>')

    # Functions
    print('## Callable')
    for attr_name, attr_obj in sorted(callable_attrs.items()):
        if not attr_name.startswith('_'):
            print('### `' + attr_name + '`')

            print('`' + str(attr_obj.__class__) + '`')
                    
            print('```')
            print(help(attr_obj))
            print('```')       
            
            print('\n<br><br>\n')


render_setup = renderSetup.instance()
callable_attrs = get_callable_attributes(render_setup)
print_markdown(callable_attrs)
```

This is what I get with Maya 2017 Update 3:

[https://gist.github.com/fredrikaverpil/510d661e4467ef4acaa0004e29c30213](https://gist.github.com/fredrikaverpil/510d661e4467ef4acaa0004e29c30213)



## Exploring Render layers

Just like how we look inside of the render setup object, we can look into a render layer object and find out what callable attributes are available, assuming you have a render layer in your scene:

```python
import maya.app.renderSetup.model.renderSetup as renderSetup

def get_callable_attributes(obj):
    """Returns the callable attributes of an object"""
    
    callable_attrs = {}
    for attr in dir(obj):
        attr_obj = getattr(obj, attr)
        if callable(attr_obj):
            callable_attrs[attr] = attr_obj
    return callable_attrs

render_setup = renderSetup.instance()
render_layers = render_setup.getRenderLayers()
render_layer = render_layers[0]
callable_attrs = get_callable_attributes(render_layer)
for callable_function in callable_attrs:
    print(callable_function)
```

This is what I get with Maya 2017 Update 3:

```
__class__
__delattr__
__format__
__getattribute__
__hash__
__init__
__new__
__reduce__
__reduce_ex__
__repr__
__setattr__
__sizeof__
__str__
__subclasshook__
__weakref__
_afterDuplicate
_afterLoadReferenceCB
_afterOpenCB
_afterUnloadReferenceCB
_beforeDuplicate
_beforeLoadReferenceCB
_beforeSaveSceneCB
_beforeUnloadReferenceCB
_cleanObservers
_decodeChildren
_decodeProperties
_encodeProperties
_getBackAttr
_getFrontAttr
_getListItemsAttr
_getNotesPlug
_hasNotesPlug
_notifyActiveLayerObservers
_onMayaNodeAddedCB
_onNodeRemoved
_preRenderLayerDelete
_switchToLayerFileIO
acceptImport
addActiveLayerObserver
addAttribute
addExternalContentForFileAttr
addListObserver
ancestors
appendChild
appendRenderLayer
attachChild
attachRenderLayer
attributeAffects
clearAll
clearListObservers
compute
connectionBroken
connectionMade
copyInternalData
createRenderLayer
creator
decode
dependsOn
detachChild
detachRenderLayer
dispose
doNotWrite
encode
forceCache
getBack
getChildren
getDefaultRenderLayer
getExternalContent
getFilesToArchive
getFront
getInternalValueInContext
getNotes
getRenderLayer
getRenderLayers
getVisibleRenderLayer
hasActiveLayerObserver
inheritAttributesFrom
initListItems
initializer
internalArrayCount
isAbstractClass
isAcceptableChild
isPassiveOutput
itemAdded
itemRemoved
legalConnection
legalDisconnection
name
parent
passThroughToMany
passThroughToOne
postConstructor
removeActiveLayerObserver
removeListObserver
setBack
setDependentsDirty
setDoNotWrite
setExternalContent
setExternalContentForFileAttr
setFront
setInternalValueInContext
setMPSafe
setNotes
shouldSave
switchToLayer
switchToLayerUsingLegacyName
thisMObject
type
typeId
typeName
```

Again, we can access the `help(function)` and see if that helps revealing useful "callables" (such as methods) by re-using the previously used `get_callable_attritbutes()` and `print_markdown()` functions (so make sure those are already sourced):

```python
import maya.app.renderSetup.model.renderSetup as renderSetup

render_setup = renderSetup.instance()
render_layers = render_setup.getRenderLayers()
render_layer = render_layers[0]
callable_attrs = get_callable_attributes(render_layer)
print_markdown(callable_attrs)
```

And this is the markdown generated, put into a gist:

[https://gist.github.com/fredrikaverpil/c7dddc44b87c8a3ee4d1007e9c240904](https://gist.github.com/fredrikaverpil/c7dddc44b87c8a3ee4d1007e9c240904)


## Exploring Render layer collections

We can use the `getCollections()` of a render layer object and look into it as well. Again, assuming you've created a collection in a render layer and have the `get_callable_attributes()` and `print_markdown()` functions sourced.

```python
import maya.app.renderSetup.model.renderSetup as renderSetup

render_setup = renderSetup.instance()
render_layers = render_setup.getRenderLayers()
render_layer = render_layers[0]
collections = render_layer.getCollections()
collection = collections[0]
callable_attrs = get_callable_attributes(collection)
print_markdown(callable_attrs)
```

Here, the markdown generated for the collection:
[https://gist.github.com/fredrikaverpil/c9e8fc601025b8fc11646c18563c853b](https://gist.github.com/fredrikaverpil/c9e8fc601025b8fc11646c18563c853b)

...and so on.


### What tricks are you using?

Please use the comments below to share any findings! :)
