let countyCode = 'US';

let phoneMaskSchema = {
    'RU': '+{7}(000) 000-00-00',
    'US': '+{1}(000) 000-0000'
};

function getCounty() {
    document.getElementById('phone').value = "";
    let county = document.getElementById('county');
    countyCode = county.options[county.selectedIndex].value;
    console.log(countyCode, phoneMaskSchema[countyCode])
}

function phoneMask() {
    let phoneNumber = document.getElementById('phone');
    let maskOptions = {
        mask: phoneMaskSchema[countyCode]
    };
    IMask(phoneNumber, maskOptions);
}

function checkPhone() {
    let phoneNumber = document.getElementById('phone');
    console.log(phoneNumber)
}

function goodResponse(name) {
    document.getElementsByClassName("overlay")[0].style.background = "#7fdfb2";
    document.getElementsByClassName("overlay")[0].style.background = "linear-gradient(to right, #48f2b1, #7fdfb2)";
    document.getElementById("status").innerHTML = "Hey!";
    document.getElementById("statusText").innerHTML = `What's up, ${name}?`;
}

function badResponse(response) {
    document.getElementsByClassName("overlay")[0].style.background = "#FF416C";
    document.getElementById("status").innerHTML = "Something went wrong..."
    document.getElementById("statusText").innerHTML = response;
}

function getUserByPhone(phone) {
    fetch(`http://0.0.0.0:8000/user/${phone}`)
        .then(response => response.json())
        .then(data => {
            if (data["error"]) {
                badResponse(data["error"])
            } else {
                goodResponse(data["name"])
            }
        })
        .catch(error => {
            badResponse(error)
        })
}

const form = document.querySelector('form')
form.addEventListener('submit', event => {
    event.preventDefault()
    let phoneNumber = document.getElementById('phone').value.replace(/\(|\)|\-|\s+/g, '')
    getUserByPhone(phoneNumber)
})