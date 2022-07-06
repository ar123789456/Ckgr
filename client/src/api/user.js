//Ручка для link

function GetAll() {
    console.log("GET ALL USER");
}

function Signup(fVal) {
    fetch('http://localhost:8080/link/new', {
        method: 'POST',
    })
        .then(response => response.json())
        .then(json => console.log(json))
}

function SignIN(fVal) {
    fetch('http://localhost:8080/link/new', {
        method: 'POST',
    })
        .then(response => response.json())
        .then(json => console.log(json))
}


function Update(fVal) {
    fetch('http://localhost:8080/link/new', {
        method: 'POST',
    })
        .then(response => response.json())
        .then(json => console.log(json))
}


function Delete(id) {
    fetch('http://localhost:8080/link/delete', {
        method: 'DELETE',
    })
        .then(response => response.json())
        .then(json => console.log(json))
}

export {
    GetAll,
    Delete,
    SignIN,
    Signup,
    Update
}