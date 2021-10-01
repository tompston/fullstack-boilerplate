import './style.css'

document.querySelector('#app').innerHTML = `
  <h2>
    <a href="http://localhost:5000/api" target="_blank"> Link </a> to api in local development, no nginx
  </h2>

  <h2>
    <a href="http://localhost/api" target="_blank"> Link </a> to api in local development, with nginx (using docker-compose.local.yml)
  </h2>
`
