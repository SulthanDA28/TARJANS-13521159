import './App.css';
import React from 'react';
import {
  Typography,
  TextField,
  Button,
  Box
} from '@mui/material';
import axios from 'axios';

import { Graphviz } from 'graphviz-react';

function App() {
  const [input, setInput] = React.useState('');
  const [graph, setGraph] = React.useState('');
  const [scc, setScc] = React.useState([]);
  const [bridge, setBridge] = React.useState([]);
  const [runtime, setRuntime] = React.useState('');
  const handleFileSelect = async (event) => {
    const file = event.target.files[0];
    const reader = new FileReader();

    reader.onload = () => {
      const text = reader.result;
      setInput(text);
    };

    reader.readAsText(file);
  };

  const ubahkeGraph = () =>{

    var arr = input.split('\n');
    const modify = input.replace(/\n/g, '\n');
    var graph = '';
    for(var i=0;i<arr.length;i++){
      if(arr[i] === '') continue;
      var temp = arr[i].split(' ');
      if(temp.length ===2){
        if(temp[0]==='' || temp[1]===''){
          graph = ''
          alert('Input tidak sesuai format');
          break
        }
      graph += temp[0] + '->' + temp[1] + ' ';
      }
      else{
        graph = ''
        alert('Input tidak sesuai format');
        break
      }
    }
    if(graph !== ''){
      axios.defaults.baseURL = 'http://localhost:8080';
      axios.post('/hasil',{
        input: modify
      }).then((res)=>{
        setBridge(res.data.hasilbridge);
        setScc(res.data.hasilscc);
        setRuntime(res.data.runtime);
      },[])
    }
    setGraph(graph);
  }

  const ubahkegraphscc = scc.map((item)=>
  <Graphviz dot={`digraph{${item.hasil}}`} options={{ width: 200, height: 200 }} />)
  // <Typography variant="h5" component="div" gutterBottom>{item.hasil}</Typography>)
  const ubahkegraphbridge = bridge.map((item)=>
  <Graphviz dot={`digraph{${item.hasil}}`} options={{ width: 200, height: 200 }} />)
  // <Typography variant="h5" component="div" gutterBottom>{item.hasil}</Typography>)



  return (
    <div>
      <Typography variant="h3" component="div" gutterBottom>
        Tarjans Algorithm
        </Typography>
      <Typography variant="h5" component="div" gutterBottom>
        Masukkan Simpul
      </Typography>
      <TextField id="outlined-basic" label="Simpul" variant="outlined" multiline
      rows={5}
      sx = {{width: '50ch'}}
      value={input}
      onChange={(e)=>setInput(e.target.value)}/>
      <input type="file" accept=".txt" onChange={handleFileSelect} />
      <Box sx={{ display: 'flex', flexDirection: 'row', p: 1, m: 1 }}>
        <Button variant="contained" sx={{ mt: 6, mb: 3 }} onClick={ubahkeGraph}>Start</Button>
      </Box>
      <Box sx={{ display: 'flex', flexDirection: 'row', p: 1, m: 1 }}>
        <Typography variant="h5" component="div" gutterBottom>
          Graph
        </Typography>
      </Box>
      <Box sx={{ display: 'flex', flexDirection: 'column', p: 1, m: 1 }}>
        <Graphviz dot={`digraph{${graph}}`} options={{ width: 200, height: 200 }} />
      </Box>
      <Box sx={{ display: 'flex', flexDirection: 'row', p: 1, m: 1 }}>
        <Typography variant="h5" component="div" gutterBottom>
          SCC
        </Typography>
      </Box>
      <Box sx={{ display: 'flex', flexDirection: 'column', p: 1, m: 1 }}>
        {ubahkegraphscc}
      </Box>
      <Box sx={{ display: 'flex', flexDirection: 'row', p: 1, m: 1 }}>
        <Typography variant="h5" component="div" gutterBottom>
          Bridge
        </Typography>
      </Box>
      <Box sx={{ display: 'flex', flexDirection: 'column', p: 1, m: 1 }}>
        {ubahkegraphbridge}
      </Box>
      <Box sx={{ display: 'flex', flexDirection: 'row', p: 1, m: 1 }}>
        <Typography variant="h5" component="div" gutterBottom>
          Runtime: {runtime}
        </Typography>
      </Box>
    </div>
  );
}

export default App;
