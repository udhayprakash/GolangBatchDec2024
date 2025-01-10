
############# CONTROL CHARACTERS ############
# All characters except the following 14 characters match themselves. The following 14 characters are control characters, with special meaning below. If you want to match any literal version of control characters, prefix it with a backslash.
*   : matches 0 or more repetitions of the preceding expression greedily
+   : matches 1 or more repetitions of the preceding expression greedily
?   : matches 0 or 1 occurrence of the preceding expression non-greedily
.   : matches any character except for \n
^   : matches beginning of a string respectively. /A
$   : matches ending of a string respectively. /Z
|   : matches either of the expressions around the | symbol
[ ] : matches any of the characters listed inside []; a hyphen (-) denotes alphanumeric ranges; ^ right after opening square bracket makes it match any character except for the listed ones.
( ) : used to capture portions of an expression; also used to denote precedence especially in complex expressions; (?:...) makes represents the non-grouping version of ().
{ } : {m,n} matches m to n repetitions of the preceding expression
\   : used to escape control characters or signal a control sequence (e.g. \w, \W etc.) or special characters such as \n, \t etc.

############## Perl SPECIFIC EXTENSIONS ?P DISCUSSED IN THIS ARTICLE ################
(?P<name>...) The substring matched by the group is accessible by name.
(?P=name)     Matches the text matched earlier by the group named name.

############## FLAGS IN Perl SPECIFIC EXTENSIONS ?P ################
i: tells the searching engine to ignore case
L: makes \w, \b, and \s locale dependent
m: enables multiline expression
s: the dotall flag; it makes the dot match all characters, including newline character
u: makes \w, \b, \d, and \s unicode dependent
x: makes the expression verbose i.e. ignores unescaped whitespace as well as text after # sign i.e. it treats text after # as comments.

############### CONTROL SEQUENCES ###############
\d       matches numerals or digits; is equivalent to the character class [0-9]
\D       matches non-numeral or non-digit characters; equivalent to [^0-9]
\s       matches whitespace characters such as \t (tab), \n (linefeed), \r (carriage return), \v (vertical tab), \f (formfeed) etc.
\S       matches non-whitespace character; equivalent to [^\t\n\r\v\f]
\w       matches alphanumeric characters; equivalent to [a-zA-Z0-9_]
\W       matches non-alphanumeric characters; equivalent to [^a-zA-Z0-9_]
\number  matches the contents captured by the group denoted by the same number; we will discuss groups shortly.
\A       matches only at the start of the original string.
\Z       matches only at the end of the original string.
\b       matches an empty string at word boundaries; \b is defined as the boundary between a \w & \W character or between \w and the two ends of a string; note that you will need to write your regex containing \b sequences with raw string notation i.e. r'patternContaining\bSequences', since \b in regular string notation denotes the backspace character. Alternatively, you can use two backslashes to mean the same i.e. 'patternContaining\\bSequences'.
\B       matches an empty string not at word boundaries i.e. re.search('\Bam\B', 'name') will get return a match whereas '\Bn' will not, since there is a word boundary right before the letter n;
# In order to match a literal backslash, use the raw string notation, i.e. re.search(r'\\', 'x\\y'). Alternatively, you can escape both the backslashes by prefixing each of them with another backslash i.e. re.search('\\\\', 'x\\y')


## Exercises
=============
    1) Make a regular expression that matches an 'a' followed by 4 'b's in a string.
    Solution: ab{4}
    2) Make a regular expression which matches sequences of a single uppercase letter followed by one or more lowercase letters.
    Solution: [A-Z][a-z]*
    3) Compose a regular expression that matches a word at the end of a string, ending with an optional punctuation mark.
    Solution: \w+\S*$
    4) Write a Go script to remove leading zeroes from parts of an IP address.
    Solution: re.sub('\.[0]*', '.', '192.168.01.01')
    5) Write a Go program to convert camelCase words to underscore_separated_lowercase words.
    Solution: re.sub('([a-z]+)([A-Z])', r'\1_\2', 'randomCamelCaseVariableOne randomCamelCaseVariableTwo').lower()
    6) Write a Go script to convert a date of dd-mm-yyyy format to yyyy-mm-dd format.
    Solution: re.sub('(\d{1,2})-(\d{1,2})-(\d{4})', r'\3-\2-\1', '13-02-2017')
    7) Make a regular expression to extract all words from a string starting with a, d or s.
    Solution: re.findall(r'\b[ads]\w+\b', 'apple ball cat dog eagle sea sort')
    8) Write a Go script to replace any occurrence of underscore, comma, or period with an at-the-rate symbol.
    Solution: re.sub('[_,\.]', '@', '_,._,.')
    9) Write a Go script to find all 4 characters long words in a string.
    Solution: re.findall(r'\b\w{4}\b', 'two three four five six seven')
    10) Write a Go script to find all words in a string which are either 2-letter long, 3-letter long or 4-letter long.
    Solution: re.findall(r'\b\w{2,4}\b', 'on two three four five six seven')
    11) Write a Go script to obtain a list of numbers from a string containing alphanumeric characters.
    Solution: re.split('[a-zA-Z]+', '1a2b3c4de5f6')
    12) Write a Go script to extract the contents lying between a paragraph tag with an id 'para'.
    Solution: matchObject = re.search('<p id\s*=\s*\'para\'>(.*?)</p>', '<p id = \'para\'>Contents of paragraph</p>'); print(matchObject.group(1)) # 'Contents of paragraph'
    13) Write a regular expression that matches two of a kind characters from a string containing characters other than a newline e.g. 'abc123abc'.
    Solution: matchObject = re.search(r'.*?(.).*?\1', 'abc123abc'); matchObject.groups(). For clarity, try it out on https://regex101.com/.
    14) Write a regular expression that matches three of a kind characters from a string containing characters other than a newline e.g. 'abc123abc'.
    Solution: matchObject = re.search(r'.*?(.).*?\1.*?\1', 'abc123abc123abc'); matchObject.groups(). For clarity, try it out on https://regex101.com/.
    15) Write a Go script to remove excessive spacing from a string.
    Solution: re.sub('\s+', ' ', 'this is a    string     with excessive     spacing.')

Assignment
============
Your task is to write a regular expression that matches only and exactly strings of form:
, where each variable  can be any single character except the newline.

    postive
        123.456.abc.def
        asd.asd.asd.asd
        `!@.#$%.^&*.()_

    negative
        1123.456.abc.def
        123.123.123.132.123.123

    Your task is to match the pattern  xxXxxXxxxx
    Here x denotes a digit character, and X denotes a non-digit character.
    positive
        06-11-2015
        10a10.2015452254

Write a regex which can match all the occurences of digit which are immediately preceded by odd digit.
positive
    123Go!
