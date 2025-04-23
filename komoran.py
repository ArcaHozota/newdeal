import sys
from konlpy.tag import Komoran

text = sys.stdin.read()
komoran = Komoran()
print(" ".join(komoran.nouns(text)))
