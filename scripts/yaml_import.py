#!/usr/bin/env python
# -*- coding: utf-8 -*-
import os
import json
import sys

import yaml
from yamlinclude import YamlIncludeConstructor
YamlIncludeConstructor.add_to_loader_class(loader_class=yaml.FullLoader, base_dir='./values/')

if __name__ == '__main__':
    with open(sys.argv[1], 'r') as f:
        data = yaml.load(f, Loader=yaml.FullLoader)
    print(yaml.dump(data))
