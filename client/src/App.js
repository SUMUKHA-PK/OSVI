import React, { Component } from 'react';
import './App.css';
import charpy from './images/charpy.jpg'
import { httpPost, httpGet } from './methods/requests';
import { error_server } from './errors';

var CameraIp = "http://210.212.194.12:8888/feed1.webm"

class View extends Component {
	constructor() {
		super()
		this.state = {
			access: true,
			infos: {
				trigger: false,
				liveStream: false,
			},
			viewers: 0,
		}
	}

	componentDidMount() {

		window.addEventListener("beforeunload", (ev) => {
			ev.preventDefault();
			httpPost('/active', { data: -1 })
				.then((r) => {

				})
			return ev.returnValue = 'Are you sure you want to close?';
		});

		httpPost('/active', { data: 1 })
			.then(r => {
				if (!r.data.Data) {
					throw new Error(error_server)
				}
			})
			.catch(e => {
				alert(error_server)
			})

		var cmp = this

		setInterval(() => {
			httpGet("/active")
				.then(r => {
					cmp.setState({
						viewers : r.data.Data
					})
					return httpGet('/trigger')
				})
				.then(r => {
					cmp.setState({
						access : Boolean(r.data.Data)
					})
				})
				.catch(e => {
					alert(error_server)
				})
		}, 100)

	}

	liveStreamClicked() {
		var infos = { ...this.state.infos }
		infos.liveStream = !infos.liveStream
		this.setState({
			infos
		})
	}

	triggerClicked() {
		httpPost("/trigger", { "RequestType": 'Start', CameraIp })
			.then(r => {
				var infos = { ...this.state.infos }
				infos.trigger = !infos.trigger
				infos.liveStream = true
				this.setState({
					infos
				})
			})
	}

	liveStream() {
		if (this.state.infos.liveStream)
			return (
				<video className="player w3-display-middle" type="text/html" width="100%" height="90%" autoPlay
					title="IET"
					src={CameraIp}
					frameBorder="0">
				</video>
			)
		else
			return (
				<span>
					<span className="w3-text-white w3-bold w3-xxlarge w3-hide-medium w3-hide-small">Charpy Experiment</span> <br />
					<img src={charpy} alt="Charpy experiment" height="70%" />
				</span>
			)
	}

	accessParameters() {
		var tName = "w3-green"
		var tText = "Trigger"

		if (this.state.access)
			return (
				<div className="w3-container" style={{ padding: "0px!important", margin: "0px!important" }}>
					<button className={'w3-button w3-margin w3-padding ' + tName} onClick={() => this.triggerClicked()}>{tText}</button>
				</div>
			)
	}

	parameters() {
		var lName = (this.state.infos.liveStream === true) ? "w3-green" : "w3-red";
		var lText = (this.state.infos.liveStream === true) ? "Streaming" : "Stream";

		return (
			<div className="w3-container w3-center">
				<button className={'w3-button w3-margin w3-padding ' + lName} onClick={() => this.liveStreamClicked()}>{lText}</button>
				{this.accessParameters()}
			</div>
		)
	}

	render() {
		return (
			<div className="w3-container">
				<h1 className="w3-center">Charpy Experiment</h1>
				<div className="w3-container w3-row w3-padding-64">
					<div className="w3-col w3-black l8 m8 s12 w3-display-container w3-padding w3-center" style={{ maxHeight: '480px', height: '60vw' }}>
						{this.liveStream()}
						<div className="w3-display-bottomleft w3-container w3-padding">
							<i className="fa fa-eye w3-text-white" aria-hidden="true"></i>
							<span className="w3-margin">{this.state.viewers}</span>
						</div>
					</div>
					<div className="w3-col l4 m4 s12 w3-container">
						<h1 className="w3-center w3-large">Controls</h1>
						{this.parameters()}
					</div>
				</div>
			</div>
		)
	}
}

function App() {
	return (
		<div>
			<View />
		</div>
	);
}

export default App;
