import sys
from konlpy.tag import Komoran

text = sys.stdin.read()
komoran = Komoran()
tokens = komoran.nouns(text)
print(" ".join(tokens))
