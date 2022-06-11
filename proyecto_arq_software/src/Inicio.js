
import React from 'react';
import './Inicio.css';

const Inicio = () => {
  const [equipo, setEquipo] = React.useState([])
  React.useEffect(()=>{
    //console.log('useEffect')
    obtenerDatos()
  }, [])
  const obtenerDatos = async()=>{
    const data = await fetch('http://localhost:8080/product')
    const users = await data.json()
    //console.log(users)
    setEquipo(users)
  }
  return (
    <div>
      <h1>Inicio</h1>
      {
          equipo.map(item => (
            <div className="grid">
            <li key = {item.id}>

      <div className="col">
          <div className="card shadow-sm">
            <img className="bd-placeholder-img card-img-top" width="auto" height="225"src={item.picture} alt="Producto" />

            <div className="card-body">
              <p className="card-text">{item.name}</p>
              <div className="d-flex justify-content-between align-items-center">
                <div className="btn-group">
                  <button type="button" className="btn btn-sm btn-outline-secondary">COMPRAR</button>
                </div>
                <p className="text-muted">{item.price}</p>
              </div>
            </div>
          </div>
        </div>
        </li>
        </div>
          ))
        }
    </div>
  )
}

export default Inicio