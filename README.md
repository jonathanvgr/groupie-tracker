# **GROUPIE TRACKER**

## Folder Structure

```
.
├── client/
│   ├── HTML / Go Templates
│   ├── js/
│   │   └── Javascript Files
│   └── style/
│       └── CSS Files
├── constants
├── Loaders/
│   └── Files needed to load data in client HTML - Paired with according route
├── Routes/
│   └── HTML navigation routes - Can be solo if no data is needed
└── main.go
```

## Server

### Routes

- routes.go
  - Initialize all routes
- \<route-name>.go
  - create route and handle how it should behave

### Loaders

- loaders.go
  - functions / variables shared by package
- \<loader-name>.go
  - should be associated with a route
  - only use if you need to load data in a page
  - always load data as : `{Data: any}`

### Constants

- constants shared by all files

### Main

- gets all the siht together

## Client

To load JS or CSS, their routes in HTML should be `/js/file.js` or `/style/file.css`. Can be changed in **main.go**

To display loaded data, use the `{{.Variable}}` syntax according to what you load via the loader file

For exemple, if you load this type in go :

```go
type RandomData struct {
    Name string
    Age uint
    IsCool bool
}

type PageData struct {
    Data RandomData
}
```

then you should do in HTML

```html
<div>{{.Data.Name}}</div>
```

It then displays as whatever you pout in Name field
