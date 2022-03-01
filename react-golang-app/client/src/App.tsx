import './App.css';
import { Header } from './components/Header/Header'
import { TodoList } from './components/TodoList/TodoList';
function App() {
  return (
    <div className="App">
      <Header />
      <TodoList todos={[
        {title: "Do dishes", isCompleted: true},
        {title: "Mow the lawn", isCompleted: false}
        ]}/>
    </div>
  );
}

export default App;
