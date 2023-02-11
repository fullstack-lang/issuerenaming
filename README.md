# issuerenaming

This repo is the use case mentioned in issue https://github.com/golang/go/issues/57559

## compiling and running the application

Requirements:
 - go (version >= 1.19)
 - node (version >= 18.xx)
 - vscode (>= 1.75)
 - gopls >= 0.37

```
git clone https://github.com/fullstack-lang/issuerenaming
cd issuerenaming
cd ng
npm i
npm install -g @angular/cli@15
ng build
cd ../go/cmd/issuerenaming
go build
./issuerenaming -unmarshallFromCode=stage.go -marshallOnCommit=stage 
```

## navigating to the application

with a navigator (for unknown reasons, it does not work with firefox on mac), navigate to http://localhost:8080/.

then navigate to the display of the UML diagram by clicking on the `Diagrams view` then select `New Diagram`.

You should see

<img width="762" alt="Screenshot 2023-02-09 at 08 14 49" src="./UML diagram before renaming.png">

## Persistance of data in go code

https://github.com/fullstack-lang/issuerenaming/blob/main/go/cmd/issuerenaming/stage.go

## Persistance of the UML diagram in go code

https://github.com/fullstack-lang/issuerenaming/blob/main/go/diagrams/NewDiagram.go

## Renaming Foo to Bar

in go/models/foo.go

```go
type Foo struct {
	Name   string
	Waldos []*Waldo
}
```

select `Foo` and rename it `Bar`

check the diffs in stage.go & NewDiagram.go
## Re generation and running the application after renaming Foo to Bar

You can regenerate the application

```
go install github.com/fullstack-lang/gong/go/cmd/gongc@issuerenaming
gongc go/models
cd go/cmd/issuerenaming
./issuerenaming -unmarshallFromCode=stage.go -marshallOnCommit=stage 
```

alternatively, you can swith to branch "afterrenaming" and rebuild the applicationo

```
cd ng
ng build
cd ../go/cmd/issuerenaming
go build
cd go/cmd/issuerenaming
./issuerenaming -unmarshallFromCode=stage.go -marshallOnCommit=stage 
```

navigate to the UML diagram.

<img width="762" alt="Screenshot 2023-02-09 at 08 14 49" src="./UML diagram after renaming.png">

Thanks to the workaround, the UML diagram has been preserved. The workaround cannot deal 
with doc link in comment, therefore both note links with "Foo" in their identifier have been lost.