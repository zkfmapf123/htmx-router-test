# HTMX / Gin test

```sh
    make run
```

- templ 은 뭔가 생각보다 더 어려울것같고, 요구사항에는 적합하지 않음
- htmx 내 hx-* 자체를 이해를 해야할듯함

```html
<div id="counter"> Counter {{.counter}} </div>
    <div id="button-container">
        <button id="btn-increment" hx-post="/counter/increment" hx-target="#counter"> Increment </button> ## 여기서 #counter 의 id 값을 innerHTML로 수정하는 거임
        <button id="btn-decrement" hx-post="/counter/decrease" hx-target="#counter"> Decrement</button>
    </div>
```

```go
c.String(http.StatusOK, fmt.Sprintf("Counter : %d", counter)) ## InnerHTML
```

## Reference

- <a href="https://templ.guide/project-structure/project-structure"> Templ </a>
- <a href="https://github.com/air-verse/air"> Hot Reload (Air) </a>
- <a href="https://htmx.org/docs/"> HTMX </a>