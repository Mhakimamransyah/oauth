
const user              = localStorage.getItem("user");
const route             = window.location.pathname;
const urlParams         = new URLSearchParams(window.location.search);
const backendUrl        = "http://127.0.0.1:8001/api"
const GithubOauthPath   = "/web/oauth/github";
const GoogleOauthPath   = "/web/oauth/google";
const Account           = {
    1 : "Github Account", 
    2 : "Google Account"
}

// home page callback
const HomePage = () => {

    const renderProfile = () => {
        const userObj =  JSON.parse(localStorage.getItem("user"));
        document.getElementsByClassName("profile-img")[0].src = userObj.image || document.getElementsByClassName("profile-img")[0].src;
        document.getElementsByClassName("profile-name")[0].innerHTML = userObj.name;
        document.getElementsByClassName("profile-email")[0].innerHTML = userObj.email;
    }

    const fetchSuccessLoginAttempt = () => {
        
        const userObj =  JSON.parse(localStorage.getItem("user"));

        fetch(backendUrl+"/users", {
            method : "GET",
            headers : {
                'content-type'  : 'application/json',
                'Authorization' : 'Bearer '+userObj.token
            },
        }).then(  response => {

            if (response.status == 401) {
                logout()
                return Promise.reject(response.statusText);
            }

            return response.json()
        }
        ).then(data => {
            
            document.getElementById("count").innerHTML = "<p><b> You have successfully logged in "+data.count_login + " Times Using "+Account[data.account]+"</b></p>";

        }).catch (error => {
            alert(error);            
        }); 
    }

    renderProfile();

    fetchSuccessLoginAttempt();

}

// logout
const logout = () => {
    localStorage.removeItem("user");
    window.history.pushState({}, "", "/web/login");
    locationHandler();
}

// page router
const routes = {
    404: {
        template: "/templates/404.html",
        title: "404",
    },
    "/web/login": {
        template: "/web/login.html",
        title: "Login Page",
    },
    "/web/home": {
        template: "/web/home.html",
        title: "Home page",
        callback : HomePage
    },
    "/web": {
        template: "/web/home.html",
        title: "Home page",
        callback : HomePage
    },
    "/web/oauth/github": {
        template: "/web/redirect.html",
        title: "Authentication"
    },
    "/web/oauth/google": {
        template: "/web/redirect.html",
        title: "Authentication"
    },
};

// router resolver
const locationHandler = async () => {

    let location = window.location.pathname; // get the url path

    // if the path length is 0, set it to primary page route
    if (location.length == 0) {
        location = "/";
    }

    // remove slash in last url path
    if (location.length > 0 && location.endsWith("/")) {
        location = location.slice(0,-1);
    }

    // get the route object from the urlRoutes object
    const route = routes[location] || routes["404"];

    // get the html from the template
    const html = await fetch(route.template).then((response) => response.text());

    // set the content of the content div to the html
    document.getElementById("content").innerHTML = html;

    // callback
    (route.callback)? route.callback() : null;

    // set the title of the document to the title of the route
    document.title = route.title;

    // set the description of the document to the description of the route
    document.querySelector('meta[name="description"]');
};

// handling oauth
if ((route == GithubOauthPath || route == GoogleOauthPath) && urlParams.get('code') != null ) {
    
    const code = urlParams.get('code');

    fetch(backendUrl+"/oauth", {

        method : "POST",
        headers : {'content-type': 'application/json'},
        body   :  JSON.stringify({
            "access_code" : code,
            "account_type" : (route == GithubOauthPath)? 1 : 2
        })

    }).then( response => {

        if (response.status != 200) {
            // error occured
            return response.json().then(data => {
                return Promise.reject(data)
            })
        } else {
            // error occured
            return response.text()
        }
    }

    ).then(data => {
    
        localStorage.setItem("user", data);
        
        window.history.pushState({}, "", "/web/home");

        locationHandler();
        
    }).catch (error => {

        alert(error.msg || "Error occured");
        
        window.history.pushState({}, "", "/web/login");

        locationHandler();

    }); 

}

if (user == null && route != "/web/login" && route != GithubOauthPath && route != GoogleOauthPath) {
    // force page to login if local storage is empty
    window.history.pushState({}, "", "/web/login");
}

locationHandler();