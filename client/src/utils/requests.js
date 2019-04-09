import store from '../redux/store/index';

var request = require('request');

//let pluginModule = require('../components/Model');
//var request = require('request');

export function save() {
    const currentProject = JSON.parse(JSON.stringify(store.getState().currProject));

    // Save current project
    fetch(`http://0.0.0.0:8000/v1/projects/`, {
        method: "POST",
        mode: "no-cors",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(currentProject)
    })
        // Log response
        .then(resp => resp.json())
        .catch(err => console.log("Error: ", err))
}

export function loadProject() {
    // Retrieve plugins from api
    fetch(`http://0.0.0.0:8000/v1/projects/${store.getState().currProject.projectName}`, {
        method: "GET",
        mode: "no-cors",
        headers: {
            "Content-Type": "application/json"
        },
    })
        // Decode response
        .then(res => { return res.json() })
        .then(data => console.log(data))
        .catch(err => console.log("Error: ", err))
}

export function loadAllProjects() {
    // Retrieve plugins from api
    fetch(`http://0.0.0.0:8000/v1/projects`, {
        method: "GET",
        mode: "no-cors",
        headers: {
            "Content-Type": "application/json"
        },
    })
        // Decode response
        .then(res => { return res.json() })
        .then(data => console.log(data))
        .catch(err => console.log("Error: ", err))
}

export function generate() {
    const currentProject = JSON.parse(JSON.stringify(store.getState().currProject));
    // Send plugins to agent
    console.log("im here")
    fetch(`http://0.0.0.0:9191/v1/templates/${store.getState().currProject.projectName}`, {
        method: "POST",
        mode: "no-cors",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(currentProject)
    })
        // Log response
        .catch(err => console.log("Error: ", err))

    let urls = `http://0.0.0.0:2379/v3alpha/watch`;


    var dataString = '{"create_request": {"key":"/vnf-agent/vpp1/config/generator/v1/template/"} }';

    var options = {
        url: urls,
        method: 'POST',
        body: dataString
    };

    function callback(error, response, body) {
        if (!error && response.statusCode === 200) {
            console.log(body);
        }
    }

    request(options, callback);
    
}

export default { save, loadProject, loadAllProjects, generate }