use std::fs::read_to_string;

#[derive(Default, Clone, Debug, Hash, PartialEq, Eq, PartialOrd, Ord)]
struct Node {
    value: String,
    left: String,
    left_index: Option<usize>,
    right: String,
    right_index: Option<usize>,
}

impl From<&String> for Node {
    fn from(value: &String) -> Self {
        let mut node: Node = Node::default();

        let split_line: Vec<String> = value.split(" = ").map(&str::to_string).collect();

        let tmp_steps: Vec<&str> = split_line[1]
        .trim()
        .trim_matches('(')
        .trim_matches(')')
        .trim_matches(char::is_whitespace)
        .split(",")
        .collect();

        node.value = split_line.first().unwrap().trim().to_string();

        node.left = tmp_steps[0].to_string();
        node.right = tmp_steps[1].trim().to_string();

        node
    }
}

impl Node {
    fn create_index(to_index: &mut Vec<Node>) {
        let search_nodes = to_index.clone();
        
        for mut node in  to_index {
            for (index, search_node) in  search_nodes.iter().enumerate() {
                if node.left == search_node.value {
                    node.left_index = Some(index);
                }
                
                if node.right == search_node.value{
                    node.right_index = Some(index);
                }
                
                if node.left_index != None && node.right_index != None{
                    break
                }
            }
            
            if node.left_index == None || node.right_index == None{
                panic!("failed to index")
            }
        }
    }
}

fn main() {
    let lines: Vec<String> =
        read_to_string("/home/feyez/coding/Christmas-2023/src/day8/src/input")
            .unwrap()
            .lines()
            .map(String::from)
            .filter(|x| !x.is_empty())
            .collect();

    let instructions: Vec<_> = lines.first().unwrap().trim().chars().collect();

    let mut nodes: Vec<Node> = lines
        .split_first()
        .unwrap()
        .1
        .iter()
        .map(Node::from)
        .collect();
    
    Node::create_index(&mut nodes);
 
    let steps = find_zzz(&nodes, &instructions);
    println!("steps: {}", steps)
}

fn find_zzz(nodes: &Vec<Node>, instructions: &Vec<char>) -> usize {
    let mut steps = 0;
    let mut current_node_index = nodes.iter().position(|node| node.value =="AAA").unwrap();
    let mut instruction_index = 0;
    
    loop {
        let instruction = instructions[instruction_index];
        
        let node_to_find = match instruction {
            'L' => (nodes[current_node_index].left_index.unwrap(), &nodes[current_node_index].left),
            'R' => (nodes[current_node_index].right_index.unwrap(), &nodes[current_node_index].right),
            _ => panic!("invalid instruction {}", instruction),
        };
        
        steps += 1;
        
        if node_to_find.1 == "ZZZ" {
            println!("done!");
            break
        }

        current_node_index = node_to_find.0;
      
        if instruction_index == instructions.len() -1{
            instruction_index = 0;
            continue;
        }
        
        instruction_index += 1;
    }

    steps
}
