<script>
	const STATES = [
		'Alabama',
		'Alaska',
		'Arizona',
		'Arkansas',
		'California',
		'Colorado',
		'Connecticut',
		'Delaware',
		'Florida',
		'Georgia',
		'Hawaii',
		'Idaho',
		'Illinois',
		'Indiana',
		'Iowa',
		'Kansas',
		'Kentucky',
		'Louisiana',
		'Maine',
		'Maryland',
		'Massachusetts',
		'Michigan',
		'Minnesota',
		'Mississippi',
		'Missouri',
		'Montana',
		'Nebraska',
		'Nevada',
		'New Hampshire',
		'New Jersey',
		'New Mexico',
		'New York',
		'North Carolina',
		'North Dakota',
		'Ohio',
		'Oklahoma',
		'Oregon',
		'Pennsylvania',
		'Rhode Island',
		'South Carolina',
		'South Dakota',
		'Tennessee',
		'Texas',
		'Utah',
		'Vermont',
		'Virginia',
		'Washington',
		'West Virginia',
		'Wisconsin',
		'Wyoming'
	];

	export let activeState;
	let selected = 'Alabama';
</script>

<div class="select-container">
	<select bind:value={selected} on:change={() => (activeState = selected)}>
		{#each STATES as state}
			<option value={state}>
				{state}
			</option>
		{/each}
	</select>
</div>

<style>
	.select-container {
		display: flex;
		justify-content: center;
		align-items: center;
		margin: auto;
		padding-top: 50%;
	}

	select {
		width: 200px;
		height: 35px;
		font-size: 18px;
	}
</style>
