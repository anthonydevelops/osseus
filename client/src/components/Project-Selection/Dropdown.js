import React from 'react';
import { Link } from 'react-router-dom';
import '../../styles/Project-Selection/Dropdown.css';

var flip = true;

class Dropdown extends React.Component {
  constructor(){
    super();
    this.state = {
      displayMenu: false,
    };
    this.showDropdownMenu = this.showDropdownMenu.bind(this);
  };

  showDropdownMenu(event) {
    flip = !flip;
    event.preventDefault();
    this.setState({ displayMenu: !flip });
  }

  setFlip() {
    flip = true;
  }

  tellMe1() {
    console.log("project 1");
    flip = true;
  }

  tellMe2() {
    console.log("project 2");
    flip = true;
  }

  tellMe3() {
    console.log("project 3");
    flip = true;
  }
  
  render() {
    return (
        <div  className="dropdown" onLoad={this.setFlip} style = {{background:"red",width:"200px"}} >
	        <div className="button" onClick={this.showDropdownMenu}> Saved Projects </div>

          { this.state.displayMenu ? (
          <ul>
    		   <li><Link onClick={this.tellMe1} to="/PluginApp">Project 1</Link></li>
    		   <li><Link onClick={this.tellMe2} to="/PluginApp">Project 2</Link></li>
    		   <li><Link onClick={this.tellMe3} to="/PluginApp">Project 3</Link></li>
          </ul>
        ):
        (
          null
        )
        }

	      </div>

    );
  }


}


export default Dropdown;