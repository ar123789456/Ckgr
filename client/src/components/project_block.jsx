import React from 'react'

export default function Project_block(props) {
  return (
    <div>
        <img src={props.url} alt={props.name} />
        <h3>{props.title}</h3>
    </div>
  )
}
