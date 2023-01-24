<script>
	import { STATE_MAP } from '../assets/constants';
	import Card, { Content, PrimaryAction, Media, MediaContent } from '@smui/card';

	export let activeState;

	let stateData;
	$: stateData = STATE_MAP[activeState];
</script>

<div class="card-display">
	<div class="card-container">
		<Card>
			<Media class="card-media-16x9" aspectRatio="16x9">
				<MediaContent>
					<!-- <h2
						class="mdc-typography--headline6"
						style="color: #fff; position: absolute; bottom: 16px; left: 16px; margin: 0;"
					>
						A card with 16x9 media.
					</h2> -->
					{#if stateData.flag !== ''}
						<img class="state-flag" src={stateData.flag} />
					{/if}
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
			<div style="padding: 1rem;">
				<h2 class="mdc-typography--headline6" style="margin: 0;">{activeState}</h2>
				<h3 class="mdc-typography--subtitle2" style="margin: 0; color: #888;">
					Length of Noncompete: {stateData.Length}
				</h3>
			</div>
			<PrimaryAction on:click={console.log('clicked')}>
				<Media class="card-media-16x9" aspectRatio="16x9" />
				<Content class="mdc-typography--body2">Literally is this even updating</Content>
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
