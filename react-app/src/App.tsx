import React from 'react';
import {
	BrowserRouter as Router,
	Switch,
	Route,
	Link
} from "react-router-dom";
import LoginScreen from './Screens/Login/LoginScreen';
import './Styles/base.scss';

function App() {
	return (
		<Router>
			<div>
				<NavigationLinks/>
				<NavigationSwitch />
			</div>
		</Router>
	);
}

function NavigationLinks(){
	return (
		<nav>
			<ul>
				<li>
					<Link to="/login">Login</Link>
				</li>
				<li>
					<Link to="/">Home</Link>
				</li>
				<li>
					<Link to="/about">About</Link>
				</li>
				<li>
					<Link to="/users">Users</Link>
				</li>
			</ul>
		</nav>
	)
}

function NavigationSwitch(){
	return (
		<Switch>
			<Route path="/login">
				<LoginScreen />
			</Route>
			<Route path="/about">
				<About />
			</Route>
			<Route path="/users">
				<Users />
			</Route>
			<Route path="/">
				<Home />
			</Route>
		</Switch>
	)
}

function Home() {
	return <h2>Home</h2>;
}

function About() {
	return <h2>About</h2>;
}

function Users() {
	return <h2>Users</h2>;
}

export default App;