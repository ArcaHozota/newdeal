import sys
from konlpy.tag import Komoran

komoran = Komoran()
tokens = komoran.morphs(sys.argv[1])
print(" ".join(tokens))
