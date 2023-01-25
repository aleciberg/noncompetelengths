<script>
	import { STATE_MAP } from '../constants';
	import Card, { Content, PrimaryAction, Media, MediaContent } from '@smui/card';

	export let activeState;

	let stateData;
	$: stateData = STATE_MAP[activeState];
</script>

<div class="card-display" style="margin-top: 60px;">
	<div class="card-container">
		<Card style="margin: 10px; display: flex; justify-content: center; align-items: center;">
			<Media class="card-media-16x9" aspectRatio="16x9">
				<MediaContent>
					<img class="state-flag" src={`state_flags/${activeState}.png`} alt="the state flag" />
				</MediaContent>
			</Media>
		</Card>
	</div>

	<!-- <div class="card-container">
		<Card style="min-width: 300px;">
			<Media class="card-media-square" aspectRatio="square">
				<div style="color: #fff; position: absolute; bottom: 16px; left: 16px;">
					<h2 class="mdc-typography--headline6" style="margin: 0;">{activeState}</h2>
					<h3 class="mdc-typography--subtitle2" style="margin: 0;">
						Length of Noncompete: {stateData.length}
					</h3>
				</div>
			</Media>
		</Card>
	</div> -->

	<div class="card-container">
		<Card>
			<div style="padding: 1rem">
				<h2 class="mdc-typography--headline6" style="margin: 0;">{activeState}</h2>
				<h3 class="mdc-typography--subtitle2" style="margin: 0; color: #888;">
					Length of Noncompete: {stateData.Length}
				</h3>
			</div>
			<PrimaryAction style=" padding: 15px" on:click={console.log('clicked')}>
				<Media class="card-media-16x9" aspectRatio="16x9" />
				<h2 class="mdc-typography--headline6" style="margin: 0; ">Our Expert Take</h2>
				<Content class="mdc-typography--body2"
					>Lorem ipsum dolor sit amet, consectetur adipiscing elit. Cras sed lacinia leo. Sed
					placerat metus nec diam congue, vitae convallis justo ultricies. Vestibulum ante ipsum
					primis in faucibus orci luctus et ultrices posuere cubilia curae; Sed lorem nisi, posuere
					quis ullamcorper id, tempor a ipsum. Mauris varius aliquet nisl ut ornare. Aliquam egestas
					lorem a est porttitor, in tincidunt arcu commodo. Nulla facilisi. Sed congue massa at
					commodo volutpat. Pellentesque cursus convallis tincidunt.</Content
				>
			</PrimaryAction>
		</Card>
	</div>
</div>

<style>
	/* * :global(.card-media-16x9) {
		background-image: url(https://place-hold.it/320x180?text=16x9&fontsize=23);
	}

	* :global(.card-media-square) {
		background-image: url(https://place-hold.it/320x320?text=square&fontsize=23);
	} */

	.state-flag {
		height: 281px;
		width: 500px;
	}
</style>
