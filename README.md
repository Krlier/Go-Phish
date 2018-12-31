# Go-Phish
Go-Phish is a software to help find possible phishing websites.

### How does it work?
The main idea is to use [DNSTwist] to generate a list of potential phishing websites from the URL provided by the user and, after that, use the concept of [Levenshtein's distance][2] to compare both website's source code and check how different they are. The website with the smaller distance calculated is then shown to the user.

If the returned percentage is very small, well, then maybe you should pay that website a visit 🧐


[DNSTwist]: https://github.com/elceef/dnstwist
[2]: https://en.wikipedia.org/wiki/Levenshtein_distance