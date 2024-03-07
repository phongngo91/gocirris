# gocirris

## How DO I MARK DOWN?


## port OUTPUT , or "WRITE" commands for CIRRIS

| do this | get this | meaning |
|---------|----------|---------|
|"EXIT" + Chr$(10)|?|go out of remote|
|||
|||


	MSComm1.Output = "PRES" + Chr$(10) -> "F"
	MSComm1.Output = "STAT" + Chr$(10) -> "("
	MSComm1.Output = "get_" + Chr$(10) -> "T"

	run the test
	MSComm1.Output = "CR_T" + Chr$(10) -> "T"



	learn
	MSComm1.Output = "LEAR" + Chr$(10) -> ""

	MSComm1.Output = "M_LE(5)" + Chr$(10) -> "F"

	MSComm1.Output = "SW_D"  is check for display button press. T - if switch is pressed

	"PUT_LIST" or "M_PUT"

	MSComm1.Output = "o_pu(CREATE TEST FROM LAST TEST SETUP)" + Chr$(10)
    MSComm1.Output = "o_put(CONNECTION RESIS 1.2 K ohm INSULATION RESIS 10 k ohm )" + Chr$(10)

	MSComm1.InBufferCount >= 1
	MSComm1.InBufferCount = 0

	Instring1 = MSComm1.Input
	STATUS1 = MSComm1.Input

[//]: # (These are reference links used in the body of this note and get stripped out when the markdown processor does its job. There is no need to format nicely because it shouldn't be seen. Thanks SO - http://stackoverflow.com/questions/4823468/store-comments-in-markdown-syntax)

   [dill]: <https://github.com/joemccann/dillinger>
   [git-repo-url]: <https://github.com/joemccann/dillinger.git>
   [john gruber]: <http://daringfireball.net>
   [df1]: <http://daringfireball.net/projects/markdown/>
   [markdown-it]: <https://github.com/markdown-it/markdown-it>
   [Ace Editor]: <http://ace.ajax.org>
   [node.js]: <http://nodejs.org>
   [Twitter Bootstrap]: <http://twitter.github.com/bootstrap/>
   [jQuery]: <http://jquery.com>
   [@tjholowaychuk]: <http://twitter.com/tjholowaychuk>
   [express]: <http://expressjs.com>
   [AngularJS]: <http://angularjs.org>
   [Gulp]: <http://gulpjs.com>

   [PlDb]: <https://github.com/joemccann/dillinger/tree/master/plugins/dropbox/README.md>
   [PlGh]: <https://github.com/joemccann/dillinger/tree/master/plugins/github/README.md>
   [PlGd]: <https://github.com/joemccann/dillinger/tree/master/plugins/googledrive/README.md>
   [PlOd]: <https://github.com/joemccann/dillinger/tree/master/plugins/onedrive/README.md>
   [PlMe]: <https://github.com/joemccann/dillinger/tree/master/plugins/medium/README.md>
   [PlGa]: <https://github.com/RahulHP/dillinger/blob/master/plugins/googleanalytics/README.md>