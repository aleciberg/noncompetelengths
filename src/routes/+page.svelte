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
		<p class="footer-text">Made with ♥ by Alec Iberg and Jordan Melvin</p>
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
		background: rgb(90, 185, 161);
		background: linear-gradient(
			90deg,
			rgba(90, 185, 161, 1) 0%,
			rgba(44, 158, 113, 1) 49%,
			rgba(21, 219, 124, 1) 100%
		);
	}

	.content-container {
		height: 80vh;
		width: 80vw;
		display: flex;
		margin-top: 100px;
		gap: 25px;
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
		color: blue;
		font-family: Montserrat;
	}

	.footer-container {
		height: 10vh;
		width: 100vw;
		display: flex;
		align-items: center;
	}
</style>
