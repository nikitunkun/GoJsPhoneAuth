# Phone Auth

Simple Phone Authorization using **GoLang** and **JavaScript**

---

## Technologies

- [![GoLang](https://img.shields.io/static/v1?label=&message=GoLang&color=blue&logo=go&logoColor=FFFFFF)](https://go.dev/)
- [![HTML](https://img.shields.io/static/v1?label=&message=HTML&color=FF9633&logo=html&logoColor=FFFFFF)](https://wikipedia.org/wiki/HTML) [![CSS](https://img.shields.io/static/v1?label=&message=CSS&color=338BFF&logo=css&logoColor=FFFFFF)](https://wikipedia.org/wiki/CSS)
- [![JavaScript](https://img.shields.io/static/v1?label=&message=JavaScript&color=black&logo=javascript&logoColor=)](https://www.javascript.com/)

---

## Preview

- Home Screen

<img src="pictures/preview.png" width="720" height="380"/>

- Choose county for the phone mask

<img src="pictures/county.png" width="720" height="380"/>

- Enter phone number. It's automaticly formatted with to the phone mask

<div>
    <img src="pictures/usa.png" width="310" height="160"/>
    <img src="pictures/russia.png" width="310" height="160"/>
</div>

- Existing Phone Number (the name is taken from API, in the example it is Nikita)

<img src="pictures/ok.png" width="720" height="380"/>

- Invalid Phone Number (the error is also taken from API)

<img src="pictures/badrequest.png" width="720" height="380"/>

---

## Using and Installation

### Instalation

- Clone or download this GitHub repository
```bash
git clone url
```
- Navigate to the project folder
```bash
cd GoJsPhoneAuth
```

### Run BackEnd

- Navigate to the backend directory
```bash
cd backend
```
- Build GoLang Project (it is simple using Makefile)
```bash
make build
```
- Run the builded project file **apiserver** in **backend** directory

### Run FrontEnd

- Navigate to the frontend directory
```bash
cd frontend
```
- Open the file **index.html** in your browser

---

## Backend Endpoints

- To create a new user
(POST method)
```http
http://0.0.0.0:8000/user
```
    Params: {
        "name": "Example",
        "phone": "+123456789",
    }

- To get the user by phone number
(GET method)
```http
http://0.0.0.0:8000/user/{phone}
```

- To delete user by phone number
(DELETE method)
```http
http://0.0.0.0:8000/user/{phone}
```

---

## Have Fun :upside_down_face:

<div>
    <img src="https://cultofthepartyparrot.com/parrots/hd/congapartyparrot.gif" width="30" height="30"/>
    <img src="https://cultofthepartyparrot.com/parrots/hd/congapartyparrot.gif" width="30" height="30"/>
    <img src="https://cultofthepartyparrot.com/parrots/hd/congapartyparrot.gif" width="30" height="30"/>
    <img src="https://cultofthepartyparrot.com/parrots/hd/congapartyparrot.gif" width="30" height="30"/>
    <img src="https://cultofthepartyparrot.com/parrots/hd/congapartyparrot.gif" width="30" height="30"/>
    <img src="https://cultofthepartyparrot.com/parrots/hd/congapartyparrot.gif" width="30" height="30"/>
    <img src="https://cultofthepartyparrot.com/parrots/hd/congapartyparrot.gif" width="30" height="30"/>
    <img src="https://cultofthepartyparrot.com/parrots/hd/congapartyparrot.gif" width="30" height="30"/>
    <img src="https://cultofthepartyparrot.com/parrots/hd/congapartyparrot.gif" width="30" height="30"/>
    <img src="https://cultofthepartyparrot.com/parrots/hd/congapartyparrot.gif" width="30" height="30"/>
    <img src="https://cultofthepartyparrot.com/parrots/hd/congapartyparrot.gif" width="30" height="30"/>
    <img src="https://cultofthepartyparrot.com/parrots/hd/congapartyparrot.gif" width="30" height="30"/>
    <img src="https://cultofthepartyparrot.com/parrots/hd/congapartyparrot.gif" width="30" height="30"/>
    <img src="https://cultofthepartyparrot.com/parrots/hd/congapartyparrot.gif" width="30" height="30"/>
    <img src="https://cultofthepartyparrot.com/parrots/hd/congapartyparrot.gif" width="30" height="30"/>
    <img src="https://cultofthepartyparrot.com/parrots/hd/congapartyparrot.gif" width="30" height="30"/>
    <img src="https://cultofthepartyparrot.com/parrots/hd/congapartyparrot.gif" width="30" height="30"/>
    <img src="https://cultofthepartyparrot.com/parrots/hd/congapartyparrot.gif" width="30" height="30"/>
    <img src="https://cultofthepartyparrot.com/parrots/hd/congapartyparrot.gif" width="30" height="30"/>
    <img src="https://cultofthepartyparrot.com/parrots/hd/congapartyparrot.gif" width="30" height="30"/>
    <img src="https://cultofthepartyparrot.com/parrots/hd/congapartyparrot.gif" width="30" height="30"/>
    <img src="https://cultofthepartyparrot.com/parrots/hd/congapartyparrot.gif" width="30" height="30"/>
    <img src="https://cultofthepartyparrot.com/parrots/hd/congapartyparrot.gif" width="30" height="30"/>
    <img src="https://cultofthepartyparrot.com/parrots/hd/congapartyparrot.gif" width="30" height="30"/>
</div>