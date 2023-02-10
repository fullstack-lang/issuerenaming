# issuerenaming

This repo is the use case mentioned in issue https://github.com/golang/go/issues/57559

## compiling and running the application

Requirements:
 - go (version >= 1.19)
 - node (version >= 18.xx)

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

## Persistance of data in go code

https://github.com/fullstack-lang/issuerenaming/blob/main/go/cmd/issuerenaming/stage.go

## Persistance of the UML diagram in go code

https://github.com/fullstack-lang/issuerenaming/blob/main/go/diagrams/NewDiagram.go

## Re generation and running the application after renaming of Foo to Bar

```
go install github.com/fullstack-lang/gong/go/cmd/gongc@issuerenaming
gongc go/models
cd go/cmd/issuerenaming
./issuerenaming -unmarshallFromCode=stage.go -marshallOnCommit=stage 
```
