const renderProductList = (products) => {
  const productsDiv = document.getElementById('products');
  productsDiv.innerHTML = '';

  products.forEach((product) => {
    const productDiv = document.createElement('div');
    productDiv.classList.add('col-3');
    productDiv.innerHTML = `
    <div class="card" style="width: 18rem;">
      <img src=${product.featured_image} class="card-img-top" alt="">
      <div class="card-body">
        <h5 class="card-title">${product.name}</h5>
        <div class="d-flex flex-row justify-content-between">
          <p class="card-text">Gender: ${product.gender}</p>
          <p class="card-text">Category: ${product.category}</p>
        </div>
        <div class="d-flex flex-row justify-content-between">
          <p class="card-text">Brand: ${product.brand}</p>
          <p class="card-text">Color: ${product.color}</p>
        </div>
        <div class="d-flex flex-row justify-content-between">
          <p class="card-text">Rating: ${product.rating}</p>
          <p class="card-text">Price: ${product.price}</p>            
        </div>
        <button type="button" class="btn btn-primary btn-block" onclick="viewDetail(${product.ID})">View Detail</button>
      </div>
    </div>
  `;
    productsDiv.append(productDiv);
  });
};

// view all
const viewAll = async () => {
  clearSearch();
  await fetch('http://localhost:3000/api/products/get')
    .then((res) => res.json())
    .then((res) => {
      renderProductList(res.data);
    });
};

// filter by gender male
const filterMale = async () => {
  clearSearch();
  await fetch('http://localhost:3000/api/products/get?gender=male')
    .then((res) => res.json())
    .then((res) => {
      renderProductList(res.data);
    });
};

// filter by gender female
const filterFemale = async () => {
  clearSearch();
  await fetch('http://localhost:3000/api/products/get?gender=female')
    .then((res) => res.json())
    .then((res) => {
      renderProductList(res.data);
    });
};

// sort price asc
const sortPriceAsc = async () => {
  clearSearch();
  await fetch('http://localhost:3000/api/products/get?sort=price_asc')
    .then((res) => res.json())
    .then((res) => {
      renderProductList(res.data);
    });
};

// sort price desc
const sortPriceDesc = async () => {
  clearSearch();
  await fetch('http://localhost:3000/api/products/get?sort=price_desc')
    .then((res) => res.json())
    .then((res) => {
      renderProductList(res.data);
    });
};

// search
const handleEnter = (e) => {
  if (e.keyCode === 13) {
    search();
  }
};

const search = async () => {
  const searchInput = document.getElementById('search-input').value;
  await fetch(`http://localhost:3000/api/products/get?search=${searchInput}`)
    .then((res) => res.json())
    .then((res) => {
      renderProductList(res.data);
    });
};

const clearSearch = async () => {
  document.getElementById('search-input').value = '';
};

// product detail
const renderProductDetail = (product) => {
  const productsDiv = document.getElementById('products');
  productsDiv.innerHTML = '';

  const productDiv = document.createElement('div');
  productDiv.classList.add('col-3');
  productDiv.innerHTML = `
      <div class="card" style="width: 18rem;">
        <img src=${product.featured_image} class="card-img-top" alt="">
        <div class="card-body">
          <h5 class="card-title">${product.name}</h5>
          <div class="d-flex flex-row justify-content-between">
            <p class="card-text">Gender: ${product.gender}</p>
            <p class="card-text">Category: ${product.category}</p>
          </div>
          <div class="d-flex flex-row justify-content-between">
            <p class="card-text">Brand: ${product.brand}</p>
            <p class="card-text">Color: ${product.color}</p>
          </div>
          <div class="d-flex flex-row justify-content-between">
            <p class="card-text">Rating: ${product.rating}</p>
            <p class="card-text">Price: ${product.price}</p>            
          </div>
        </div>
      </div>
    `;

  productsDiv.append(productDiv);
};

const viewDetail = async (id) => {
  await fetch(`http://localhost:3000/api/products/getone?id=${id}`)
    .then((res) => res.json())
    .then((res) => {
      renderProductDetail(res.data);
    });
};
