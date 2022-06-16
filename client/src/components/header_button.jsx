import React from 'react'
import { Link } from "react-router-dom";

export default function HeaderButton (props) {
  console.log(props);
  return (
    <Link to={props.url}>{props.name}</Link>
    // <form action="google.com" className='button-form'>
    //     <button type='submit' className='button'>{props.name}</button>
    // </form>
  )
}
