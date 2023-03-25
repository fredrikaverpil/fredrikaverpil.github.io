---
date: 2013-10-11
tags:
- python
---

# Catching string from stdout with Python

Using a Python “wrapper” script you can catch the output of an executing application, which is very useful when you want to perform certain tasks depending on what is being printed to stdout by the application.

<!-- more -->


This is a simplified example of how I do it:

```python
# Imports
import os, sys, subprocess

# Build command
command = [ 'python', os.join.path('/path/to', 'scriptFile.py') ]

# Execute command
p = subprocess.Popen(command, stdout=subprocess.PIPE, stderr=subprocess.STDOUT)

# Read stdout and print each new line
sys.stdout.flush()
for line in iter(p.stdout.readline, b''):

	# Print line
	sys.stdout.flush()
	print(">>> " + line.rstrip())


	# Look for the string 'Render done' in the stdout output
	if 'Render done' in line.rstrip():

		# Write something to stdout
		sys.stdout.write('Nice job on completing the render, I am executing myFunction()\n' )
		sys.stdout.flush()

		# Execute something
		myFunction()
```


### Catching stuff from stdout for Pixar's Tractor

Here’s an example of catching the output of a V-Ray for Maya render and outputting `TR_PROGRESS nnn%` to stdout whenever a percentage is being printed to stdout, which makes [Pixar’s Tractor](http://renderman.pixar.com/view/pixars-tractor) show a task progress in the task node tree.

```python
# Imports
import os, sys, subprocess, re

# Build command
command = ['python', os.join.path('/path/to', 'mayaWrapper.py'), '-r', 'vray']

# Execute command
p = subprocess.Popen( command, stdout=subprocess.PIPE, stderr=subprocess.STDOUT)

# Read stdout and print each new line
sys.stdout.flush()
for line in iter(p.stdout.readline, b''):
	sys.stdout.flush()
	print(">>> " + line.rstrip())

	# Tractor Progress
	match = re.search('[\d]* %', line.rstrip() )			# Detect percentage without period
	if not match:
		match = re.search('[\d.\d]* %', line.rstrip() )	    # Detect percentage with period
	if match:
		percent = match.group(0).replace('%','').replace(' ','')
		try:
			percent = float(percent)
			percent = int(percent)
			percent = str(percent)
		except:
			pass
		sys.stdout.write('TR_PROGRESS ' + str(percent) + '%\n' )
		sys.stdout.flush()
```