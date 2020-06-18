# html page sample parser

# Clone
    $ git clone git@github.com:andyklimenko/html-element-finder.git

# Build
    $ go build -o html-parser main.go
    
# Help 
> -id - an id of element to look for in original html
>
> -original - relative path to original html page
>
> -updated - relative path to updated html
# Run
    $ html-parser -id=make-everything-ok-button -original=test-data/sample-0-origin.html -updated=test-data/sample-1-evil-gemini.html
    $ html-parser -id=make-everything-ok-button -original=test-data/sample-0-origin.html -updated=test-data/sample-2-container-and-clone.html
    $ html-parser -id=make-everything-ok-button -original=test-data/sample-0-origin.html -updated=test-data/sample-3-the-escape.html
    $ html-parser -id=make-everything-ok-button -original=test-data/sample-0-origin.html -updated=test-data/sample-4-the-mash.html
    
 