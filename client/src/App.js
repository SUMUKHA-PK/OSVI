import React, { Component } from 'react';
import './App.css';
import { httpGet } from './methods/requests';

class View extends Component {
	constructor() {
		super()
		this.state = {
			access: true,
			infos: {
				trigger: true
			}
		}
	}

	triggerClicked() {
		httpGet('/trigger')
		this.setState({
			infos: {
				trigger: !this.state.infos.trigger
			}
		})
	}


	parameters() {
		var tName = (this.state.infos.trigger === true) ? "w3-green" : 'w3-red';
		var tText = (this.state.infos.trigger === true) ? "ON" : 'OFF';

		if (this.state.access)
			return (
				<div className="w3-container">
					<p className="w3-center">Trigger is <button className={'w3-button w3-margin w3-padding ' + tName} onClick={() => this.triggerClicked()}>{tText}</button></p>
				</div>
			)
	}

	render() {
		return (
			<div className="w3-container">
				<h1 className="w3-center">Charpy Experiment</h1>
				<div className="w3-container w3-row w3-padding-64">
					<div className="w3-col w3-black l8 m8 s12 w3-display-container w3-padding" style={{ maxHeight: '480px', height: '60vw' }}>
						<div className="w3-display-bottomleft w3-container w3-padding">
							<i class="fa fa-eye w3-text-white" aria-hidden="true"></i>
						</div>
						<iframe className="player w3-display-middle" type="text/html" width="100%" height="90%"
							title="IET"
							src="http://210.212.194.12:8888/feed1.webm"
							frameBorder="0">
						</iframe>
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
