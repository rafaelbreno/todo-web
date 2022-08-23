<script lang="ts">
  import { onMount } from 'svelte';
  type List = {
    id: string,
    name: string;
  };

  let listForm: List = {};

  let allLists: List[] = [];

  onMount(async () => {
    const resp = await fetch('http://localhost:5001/list', {
      method: 'GET',
      headers: {'Content-Type': 'application/json'},
    });

    const json = await resp.json();
    
    if (resp.status != 200) {
      console.log(json);
    } else {
      let ls: List[] = json;
      allLists = [...allLists, ...ls];
      console.log(allLists);
    }
  })

  async function createList() {
    const resp = await fetch('http://localhost:5001/list', {
      method: 'POST',
      headers: {'Content-Type': 'application/json'},
      body: JSON.stringify(listForm),
    });

    const json = await resp.json();

    if (resp.status != 200) {
      console.log(json);
    } else {
      let l: List = json;
      allLists = [...allLists, l];
      console.log(allLists);
    }
  }

</script>

<div>
  <label>Name:</label>
  <input type="text" bind:value={listForm.name}/>
  <button on:click={createList}>Create</button>
  <ul>
    {#each allLists as list}
      <li>{list.id} | {list.name}</li>
    {/each}
  </ul>
</div>
