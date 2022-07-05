
export function GetAll() {
    fetch('http://localhost:8080/link', {
        method: 'GET',
    })
        .then(response => response.json())
        .then(json => console.log(json))
}