# IITK COIN

## ABOUT
The project aims to build a pseudo-coin system for use in the IITK Campus. We are aiming to develop the back-end system and (if time permits) a front-end app for this pseudo-currency.

## Specifications:
[IITK-COIN Specs](IITK_COIN.md)  
This is the original project proposal drafted by the SnT Council: 

## Organization

### Server Handlers:
[Different Server Endpoints](initServer.go)  
[User Authentication](authServer.go)  
[Coin related Requests to Server](coinServer.go)  
### Database Handlers:
[User Info from Database](userDB.go)  
[Display Current Entries in Database](displayDB.go)  
[Coin related Requests to Database](coinDB.go)  
Database File : students.db with Three Tables: students, transfers, awards

### Main
[Entry Point](main.go)

### Others
[go.mod](go.mod) | [go.sum](go.sum) | [IITK_COIN.md](IITK_COIN.md) | [test.md](test.md) | [.gitignore](.gitignore)
## Relavant Links
[Mid Term Presentation](https://docs.google.com/presentation/d/1oAtiYqLpoEe39rXtIOAqFTXmDqcdQXySyao5ijLXzyY/edit?usp=sharing)  
[Mid Term Documentation](https://docs.google.com/document/d/1VCocS_ilBZw1ROmFcFnpqHg7HmIL75F3cwHC7BimzzU/edit?usp=sharing)

## Usage and Testing
[Sample Client and Server Log](test.md)
