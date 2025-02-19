#  Copyright (c) 2020 Siemens AG
#
#  Permission is hereby granted, free of charge, to any person obtaining a copy of
#  this software and associated documentation files (the "Software"), to deal in
#  the Software without restriction, including without limitation the rights to
#  use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
#  the Software, and to permit persons to whom the Software is furnished to do so,
#  subject to the following conditions:
#
#  The above copyright notice and this permission notice shall be included in all
#  copies or substantial portions of the Software.
#
#  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
#  IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
#  FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
#  COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
#  IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
#  CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
#
#  Author(s): Jonas Plum

import contextlib
import importlib
import io
import os
import shutil
import sys
import tempfile

import pytest

sys.path.append("scripts/scripts")
hotfixes = importlib.import_module("elementary-hotfixes")


@pytest.fixture
def data():
    tmpdir = tempfile.mkdtemp()
    shutil.copytree(os.path.join("test", "data"), os.path.join(tmpdir, "data"))
    return os.path.join(tmpdir, "data")


def test_hotfixes(data):
    with io.StringIO() as buf, contextlib.redirect_stdout(buf):
        hotfixes.main(os.path.join(data, "example1.forensicstore"))
        lines = buf.getvalue().split("\n")
        assert len(lines) == 14 + 2

    shutil.rmtree(data)
