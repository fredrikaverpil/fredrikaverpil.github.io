---
title: File copy progress window with PyQt/PySide and shutil
tags: [python, pyside, pyqt]
draft: false

# PaperMod
ShowToc: false
TocOpen: false

created: 2015-05-12T02:00:12+02:00
updated: 2022-11-15T22:29:16+01:00
---

I recently wanted to show the progress of a file copy made from a python script (or actually, from a python script running inside of Maya and Nuke). First I looked to piggyback on the OS native ways of copying a file, but oddly enough it turned out there was no way of showing a progress bar when doing this.



So I turned to Python’s [shutil](https://docs.python.org/2/library/shutil.html) and [PySide](https://wiki.qt.io/Category:LanguageBindings::PySide) (but [PyQt](http://www.riverbankcomputing.co.uk/software/pyqt/intro) would work just as well for this). Shutil cannot show progress by itself, but when asking around about this, I [learned](http://stackoverflow.com/questions/29967487/get-progress-back-from-shutil-file-copy-thread) that you can implement your own copy function in Python using its built-in `copyfileobj`.

Below is a simple python script I came up with, to illustrate how to implement your own file copy with progress in Python.

```python
# -*- coding: utf-8 -*-
#
#
#
#	   ___      ___                                                                                                        
#	 /'___\ __ /\_ \                                                                                                       
#	/\ \__//\_\\//\ \      __    ___    ___   _____   __  __      _____   _ __   ___      __   _ __    __    ____    ____  
#	\ \ ,__\/\ \ \ \ \   /'__`\ /'___\ / __`\/\ '__`\/\ \/\ \    /\ '__`\/\`'__\/ __`\  /'_ `\/\`'__\/'__`\ /',__\  /',__\
#	 \ \ \_/\ \ \ \_\ \_/\  __//\ \__//\ \L\ \ \ \L\ \ \ \_\ \   \ \ \L\ \ \ \//\ \L\ \/\ \L\ \ \ \//\  __//\__, `\/\__, `\
#	  \ \_\  \ \_\/\____\ \____\ \____\ \____/\ \ ,__/\/`____ \   \ \ ,__/\ \_\\ \____/\ \____ \ \_\\ \____\/\____/\/\____/
#	   \/_/   \/_/\/____/\/____/\/____/\/___/  \ \ \/  `/___/> \   \ \ \/  \/_/ \/___/  \/___L\ \/_/ \/____/\/___/  \/___/
#	                                            \ \_\     /\___/    \ \_\                 /\____/                          
#	                                             \/_/     \/__/      \/_/                 \_/__/                           
#
#
# by Fredrik Averpil, http://fredrik.averpil.com, fredrik.averpil [at] gmail.com


import sys, os, datetime
from PySide import QtGui, QtCore
from optparse import OptionParser



class FileCopyProgress( QtGui.QWidget ):
	'''Custom shutil file copy with progress'''
	def __init__(self, parent=None, src=None, dest=None):
		super(FileCopyProgress, self).__init__()

		self.src = src
		self.dest = dest
		self.build_ui()



	def build_ui(self):

		hbox = QtGui.QVBoxLayout()

		lbl_src = QtGui.QLabel('Source: ' + self.src)
		lbl_dest = QtGui.QLabel('Destination: ' + self.dest)
		self.pb = QtGui.QProgressBar()

		self.pb.setMinimum(0)
		self.pb.setMaximum(100)
		self.pb.setValue(0)


		hbox.addWidget(lbl_src)
		hbox.addWidget(lbl_dest)
		hbox.addWidget(self.pb)
		self.setLayout(hbox)

		self.setWindowTitle('File copy')

		self.auto_start_timer = QtCore.QTimer()
		self.auto_start_timer.timeout.connect( lambda : self.copyfileobj( src=self.src, dst=self.dest, callback_progress=self.progress, callback_copydone=self.copydone )  )
		self.auto_start_timer.start(2000)

		self.show()


	def progress(self, fsrc, fdst, copied):
		size_src = os.stat( fsrc.name ).st_size
		size_dst = os.stat( fdst.name ).st_size

		float_src = float( size_src )
		float_dst = float( size_dst )

		percentage = int(float_dst/float_src*100)

		try:
			self.pb.setValue( percentage )
		except:
			pass

		app.processEvents()


	def copydone(self, fsrc, fdst, copied):
		self.pb.setValue( 100 )
		self.close()


	def copyfileobj(self, src, dst, callback_progress, callback_copydone, length=8*1024):

		# Prevent progress callback from being made each cycle
		c = 0
		c_max = 50

		try:
			self.auto_start_timer.stop()
		except:
			print 'Error: could not stop QTimer'


		with open(src, 'r') as fsrc:
			with open(dst, 'w') as fdst:
				copied = 0
				while True:
					buf = fsrc.read(length)
					if not buf:
						break
					fdst.write(buf)
					copied += len(buf)
					c += 1
					if c == c_max:
						callback_progress(fsrc=fsrc, fdst=fdst, copied=copied)
						c = 0
				callback_copydone(fsrc=fsrc, fdst=fdst, copied=copied)


if __name__ == '__main__':


	parser = OptionParser()
	parser.add_option("-s", "--src", dest="src",
					  help="source FILE", metavar="FILE")
	parser.add_option("-d", "--dest", dest="dest",
		 				help="destination FILE", metavar="FILE")
	(options, args) = parser.parse_args()

	if os.path.isfile( options.src ):
		app = QtGui.QApplication(sys.argv)
		ex = FileCopyProgress(src=options.src, dest=options.dest)
		sys.exit(app.exec_())
```


You can run the file like this:

```python
python filecopyprogress.py --src=c:\hello1.png --dest=c:\hello2.png
```

Please, go ahead and do whatever you wish with this code. And let me know if you improve it, such as adding support for folders!
