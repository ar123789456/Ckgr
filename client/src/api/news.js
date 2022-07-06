//Ручка для news

function GetAllForClient() {
    fetch('http://localhost:8080/link', {
        method: 'GET',
    })
        .then(response => response.json())
        .then(json => console.log(json))
}

function GetSingle(id) {
    fetch('http://localhost:8080/link', {
        method: 'GET',
    })
        .then(response => response.json())
        .then(json => console.log(json))
}

function GetAllForAdmin() {
    fetch('http://localhost:8080/link', {
        method: 'GET',
    })
        .then(response => response.json())
        .then(json => console.log(json))
}

function Update() {
    fetch('http://localhost:8080/link', {
        method: 'GET',
    })
        .then(response => response.json())
        .then(json => console.log(json))
}


function Create(fVal) {
    fetch('http://localhost:8080/link/new', {
        method: 'POST',
    })
        .then(response => response.json())
        .then(json => console.log(json))}

function Delete(id) {
    fetch('http://localhost:8080/link/delete', {
        method: 'DELETE',
    })
        .then(response => response.json())
        .then(json => console.log(json))}

export {
    GetAllForAdmin,
    GetAllForClient,
    GetSingle,
    Update,
    Create,
    Delete,    
}