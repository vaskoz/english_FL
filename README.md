# english_FL
What if you randomly shuffled the letters between the first and last in the English language. How hard would it be to descramble it while reading?

# It's pretty easy, but why?

There are some assumptions in this problem:

1. Only grammatically and logically correct English sentences are allowed.
2. The words used in those sentence require the following:
    1. The first and last letter of the word are in their correct positions.
    2. There are no substitution errors (typos) for any letters.
    3. There are no addition or removal of letters in the word.

## Okay, but why is it easy?

There aren't that many choices for possible words that meet these constraints. I wrote a program to show that **at most 4 words** could be candidates for these scrambled words.

Notice that when you review possibilities in scrambled words, only 1 is a common word and the others are increasingly exotic.

The other thing to notice is that number of unique first & last pair combinations in this dictionary is **646**. The total possible in English for a 26-letter alphabet is exactly **26^2=676**. This means there are exactly **676-646 = 30** combinations of a first letter & last letter that NEVER appear in the English dictionary.

# A program to demonstrate that there aren't many possibilities with these constraints.

```shell
go run main.go < words_alpha.txt
Unique first/last pairs 646
4
26
[ceratin certain citrean creatin]
[balistes bastiles bestials blasties]
[tairge tirage triage trigae]
[caliver caviler claiver clavier]
[denudes dudeens duendes dundees]
[saintless saltiness slatiness stainless]
[sipers speirs spiers spires]
[spirts sprits stirps strips]
[scepters sceptres specters spectres]
[slipes speils spiels spiles]
[benote betone bonete bontee]
[weird wierd wired wried]
[petrous pouters proetus proteus]
[pastries piasters piastres piratess]
[painters pantries parentis pertains]
[palets petals plates pleats]
[partless persalts plasters psalters]
[partley peartly platery prelaty]
[salter slater staler stelar]
[unparsed unrasped unspared unspread]
[warstles wartless wastrels wrastles]
[carlos carols claros corals]
[carets cartes caters crates]
[palster pastler plaster psalter]
[tares tears teras treas]
[felid field filed flied]
```

# Picture of example text

![Screenshot of scrambled English text where the first and last letter stay in the correct position, but the middle is randomly scrambled.](ScrambledEnglishReadable.jpg)