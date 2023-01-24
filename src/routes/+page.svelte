<script>
	import StateInfo from '../components/StateInfo.svelte';
	import StatePicker from '../components/StatePicker.svelte';

	let activeState = 'Alabama';
	$: console.log(activeState);

	console.log('booting up');
</script>

<div class="main-container">
	<div class="left-container">
		<StatePicker bind:activeState />
	</div>
	<div class="right-container">
		<StateInfo bind:activeState />
	</div>
</div>

<style>
	.main-container {
		border: 1px solid red;
		height: 100vh;
		width: 100vw;
		overflow: hidden;
		display: flex;
	}

	.left-container {
		border: 1px solid blue;
		height: 100vh;
		width: 50vw;
	}

	.right-container {
		border: 1px solid green;
		height: 100vh;
		width: 50vw;
	}
</style>
