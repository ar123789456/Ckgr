import React from "react";
import HeaderButton from "./header_button.jsx";
import CGALogo from "./CGA.svg"
import { Link } from "react-router-dom";

export default function Header() {
    return (
        <div className='Header' >
            <Link to="/"><img src={CGALogo} alt="logo" className="logo"></img></Link>

            
            <div className='header-button'>
            <nav>
            <HeaderButton name="Projects" url="projects"/>
            <HeaderButton name="People" url="people"/>
            <HeaderButton name="News" url="news"/>
            <HeaderButton name="Contacts" url="contact"/>
            </nav>
            </div>
            
        </div>
    )
}