import React, { useState } from "react";
import "./css/User.css";
import logo from "./images/logo.svg"
import cart from "./images/cart.svg"
import loadinggif from "./images/loading.gif"
import usersvg from "./images/user.svg"
import Cookies from "universal-cookie";

const Cookie = new Cookies();

async function getUserById(id){
    return await fetch('http://localhost:8080/user/' + id, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json'
    }
}).then(response => response.json())

}

async function getOrdersByUserId(id){
    return await fetch('http://localhost:8080/orders/' + id, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json'
    }
}).then(response => response.json())

}

async function getAddressesByUserId(id){
    return await fetch('http://localhost:8080/addresses/' + id, {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json'
    }
}).then(response => {
  if(response.status == 200){
    return response.json()
  }else{
    return []
  }
})

}

function goto(path){
  window.location = window.location.origin + path
}


function logout(){
  Cookie.set("user_id", -1, {path: "/"})
  document.location.reload()
}

function showDetails(id){
  let div = document.getElementsByClassName("o" + id)
  for(let detail of div){
    let show = false;
    detail.classList.value.split(" ").forEach(c=>{
      if (c == "hidden"){
        show = true;
      }
    })
    if(show){
      detail.classList.remove("hidden")
    }else{
      detail.classList.add("hidden")
    }
  }


}

function getDetails(details){
  return details.map((detail) =>
      <tbody>
        <tr className="detail">
          <td className="id">{detail.product_id}</td>
          <td className="name">{detail.name}</td>
          <td className="quantity">{detail.quantity}</td>
          <td className="price">{detail.currency_id + "$" + detail.price}</td>
        </tr>
      </tbody>
  )
}

function showAddressDetails(id){
  let div = document.getElementsByClassName("a" + id)
  for(let detail of div){
    let show = false;
    detail.classList.value.split(" ").forEach(c=>{
      if (c == "hidden"){
        show = true;
      }
    })
    if(show){
      detail.classList.remove("hidden")
    }else{
      detail.classList.add("hidden")
    }
  }
}

function showUserAddresses(addresses){
  return addresses.map((address) =>

  <div obj={address} key={address.address_id} className="userAddress" onClick={()=>showAddressDetails(address.address_id)}>
    <a className="addressn">{address.street1} {address.number}</a>
    <div className={"details" + " a" + address.address_id + " hidden"}>
      <div><span className="userAddressInfo"> Street: </span> <a className="userAddressInfoLoad">{address.street1}</a> </div>
      <div><span className="userAddressInfo"> Street2: </span> <a className="userAddressInfoLoad">{address.street2} </a> </div>
      <div><span className="userAddressInfo"> Number: </span> <a className="userAddressInfoLoad">{address.number} </a> </div>
      <div><span className="userAddressInfo"> District: </span> <a className="userAddressInfoLoad">{address.district} </a> </div>
      <div><span className="userAddressInfo"> City: </span> <a className="userAddressInfoLoad">{address.city} </a> </div>
      <div><span className="userAddressInfo"> Country: </span> <a className="userAddressInfoLoad">{address.country} </a> </div>
    </div>
  </div>
 )
}



function showUserOrders(orders){
  return orders.map((order) =>

  <div obj={order} key={order.order_id} className="order" onClick={()=>showDetails(order.order_id)}>
    <a className="ordern">Order N°: {order.order_id}</a>
    <a className="date">Date: {order.date.split("T")[0]}</a>
    <a className="total">Total: <span>{order.currency_id + "$" + order.total} </span></a>

    <div className={"details" + " o" + order.order_id + " hidden"}>
      <table>
        <thead>
          <tr className="detail thead">
            <td className="id">ID</td>
            <td className="name">NAME</td>
            <td className="quantity">n</td>
            <td className="price">PRICE</td>
          </tr>
        </thead>

      {getDetails(order.details)}
      </table>
    </div>

  </div>
 )
}



function User() {
  const [isLogged, setIsLogged] = useState(false)
  const [user, setUser] = useState({})
  const [cartItems, setCartItems] = useState("")
  const [orders, setOrders] = useState([])
  const [addresses, setAddresses] = useState([])


  if (Cookie.get("user_id") > -1 && !isLogged){
    getUserById(Cookie.get("user_id")).then(response => setUser(response))
    setIsLogged(true)
    if (!cartItems && Cookie.get("cartItems")){
      setCartItems(Cookie.get("cartItems"))
    }

    if(Cookie.get("orders") == undefined){
      getOrdersByUserId(Cookie.get("user_id")).then(response => {setOrders(response)})
      getAddressesByUserId(Cookie.get("user_id")).then(response => (setAddresses(response)))
    }


  }

  if (!(Cookie.get("user_id") > -1) && isLogged){
    setIsLogged(false)
  }

  const login = (

    <span>
    <img src={usersvg} onClick={()=>goto("/user")} id="user" width="48px" height="48px"/>
    <img src={cart} onClick={()=>goto("/cart ")} id="cart" width="48px" height="48px"/>
    <span className="cartNumber">{cartItems > 0 ? cartItems : 0}</span>
    <a id="logout" onClick={logout}> <span> Welcome in {user.first_name} </span> </a>
    </span>
  )

  const showUserInfo = (
    <div className="userInfo">
      <img src={usersvg} width="128px" height="128px"/>
      <div> {user.first_name} {user.last_name} </div>
      <div> Username: {user.username} </div>
      <div> Email: {user.email} </div>
      <div> Addresses: {addresses.length > 0 ? showUserAddresses(addresses) : <a> No addresses :( </a>} </div>
      <div> Orders: {orders.length > 0 ? showUserOrders(orders) : <a> No orders :( </a>}</div>
    </div>
  )

  const pleaseLogin = (
    <div> Nothing to show. Please login and maybe we'll get some info for ya </div>
  )


  const loading = (
    <img id="loading" src={loadinggif}/>
  )

  return (
    <div className="home">
      <div className="topnav">
        <div>
          <img src={logo} width="80px" height="80px" id="logo" onClick={()=>goto("/")} /> <p>3 Random Words Shop</p>
        </div>
        {isLogged ? login : <a id="login" onClick={()=>goto("/login")}>Login</a>}
      </div>

      <div id="main">
        {isLogged ? showUserInfo : pleaseLogin}
      </div>
    </div>
  );
}

export default User;
