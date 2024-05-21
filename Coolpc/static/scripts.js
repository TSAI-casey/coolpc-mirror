document.addEventListener('DOMContentLoaded', function() {
    let dropdownButtons = document.querySelectorAll('.dropdown-btn');
    dropdownButtons.forEach(button => {
        button.addEventListener('click', function() {
            this.nextElementSibling.classList.toggle('show');
        });
    });
});

let cart = [];

function searchItems() {
    let searchValue = document.getElementById('search').value.toLowerCase();
    let items = document.querySelectorAll('.item');
    items.forEach(item => {
        if (item.innerText.toLowerCase().includes(searchValue)) {
            item.style.display = 'block';
        } else {
            item.style.display = 'none';
        }
    });
}

function addToCart(name, price) {
    cart.push({ name, price });
    renderCart();
}

function removeFromCart(index) {
    cart.splice(index, 1);
    renderCart();
}

function renderCart() {
    let cartDiv = document.getElementById('cart');
    cartDiv.innerHTML = '';
    cart.forEach((item, index) => {
        let cartItemDiv = document.createElement('div');
        cartItemDiv.className = 'cart-item';
        cartItemDiv.innerHTML = `
            <span>${item.name} - $${item.price}</span>
            <button onclick="removeFromCart(${index})">Delete</button>
        `;
        cartDiv.appendChild(cartItemDiv);
    });
}
