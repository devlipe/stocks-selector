# Stocks Selector

This is a simple project written in go. It's a stocks selector that given a list of stocks, each one with their data, can filter and rank them. The project utilizes Joel GreenBlatt's Magic Formula, along with other indicators to list the stocks with the most value within.


## Deployment

To deploy this project run

```bash
  docker compose build
```
This will build the service. Since the project uses the stdin, it cannot be deployed by `docker compose up`.

To run the service

```bash
  docker compose run selector
```

This command will run it.

You can also run `go run` or `go build && ./stocks-selector` to start the app.
## Installation

As a go project, you can install it using the following commands

```bash
  go build 
  go install
```

Then you should have the following command available on your sheel


```bash
  stocks-selector
```
    
## Acknowledgements

 - [Joel Greenblatt Magic Formula](https://www.magicformulainvesting.com/)
 - [Clube do Valor - Ramiro](https://clubedovalor.com.br/)


## Authors

- [@devlipe](https://www.github.com/devlipe)


## Lessons Learned

The purpose of this project was to learn a bit more about docker, environmental variables, and go. It uses the latest golang feature, which is generics functions. It was nice to also learn how to manipulate text on the terminal, and gave the app different behaviors based on the os.
