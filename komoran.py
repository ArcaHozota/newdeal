import sys
from konlpy.tag import Komoran

text = sys.stdin.read()
komoran = Komoran()
tokens = komoran.morphs(text)
print(" ".join(tokens))
