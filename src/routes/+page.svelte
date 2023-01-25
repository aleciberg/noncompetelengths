<script lang="ts">
	import StateInfo from '../components/StateInfo.svelte';
	import StatePicker from '../components/StatePicker.svelte';

	let activeState = 'Alabama';
	$: console.log(activeState);

	console.log('booting up');
</script>

<svelte:head>
	<link rel="preconnect" href="https://fonts.googleapis.com" />
	<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin />
	<link
		href="https://fonts.googleapis.com/css2?family=Montserrat:ital,wght@1,300&display=swap"
		rel="stylesheet"
	/>
</svelte:head>

<div class="main-container">
	<div class="content-container">
		<div class="left-container">
			<StatePicker bind:activeState />
		</div>
		<div class="right-container">
			<StateInfo bind:activeState />
		</div>
	</div>
	<div class="footer-container">
		<p class="footer-text">Made with love by Alec Iberg and Jordan Melvin</p>
	</div>
</div>

<style>
	.main-container {
		height: 100vh;
		width: 100vw;
		overflow: hidden;
		display: flex;
		flex-wrap: wrap;
		flex-direction: row;
		justify-content: center;
		background-color: #7f6239;
	}

	.content-container {
		height: 80vh;
		width: 80vw;
		display: flex;
		margin-top: 100px;
		background-color: #c4a77d;
		border-radius: 20px;
	}

	.left-container {
		height: 60vh;
		width: 30vw;
	}

	.right-container {
		height: 60vh;
		width: 60vw;
	}

	.footer-text {
		font-size: 14px;
		padding-left: 10px;
		color: #2ba54f;
		font-family: Montserrat;
	}

	.footer-container {
		height: 10vh;
		width: 100vw;
		display: flex;
		align-items: center;
	}
</style>
