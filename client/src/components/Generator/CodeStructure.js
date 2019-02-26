import React from 'react';
import "../../styles/App.css";
import "../../styles/Generator/GeneratorApp.css";

/*This component represents the left webpage division where the
* plugins will reside initially. The plugins are being represented
* by the incoming prop.children. Presumedly in the future the 
* prop children can be stored in an array instead of initializing
* one-by-one.
*/
const CodeStructure = (props) => {
    return (
        <div className="body">
            <div className="split left">
            <p className="whitetextgen">Code Structure</p>
            </div>
        </div>
    );
};

export default CodeStructure;