use std::fs::read_to_string;

#[derive(Default, Clone, Debug, Hash, PartialEq, Eq, PartialOrd, Ord)]
struct Node {
    index: Option<usize>,
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

        for (base_index, node) in to_index.iter_mut().enumerate() {
            node.index = Some(base_index);

            for (index, search_node) in search_nodes.iter().enumerate() {
                if node.left == search_node.value {
                    node.left_index = Some(index);
                }

                if node.right == search_node.value {
                    node.right_index = Some(index);
                }

                if node.left_index.is_some() && node.right_index.is_some() {
                    break;
                }
            }

            if node.left_index.is_none() || node.right_index.is_none() {
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

    let steps = find_zzz(nodes, instructions);

    println!("steps: {}", steps)
}

fn find_zzz(nodes: Vec<Node>, instructions: Vec<char>) -> usize {
    let mut steps = 0;
    let mut instruction_index = 0;
    let mut current_nodes: Vec<Node> = nodes
        .iter()
        .filter(|node| node.value.ends_with('A')).cloned()
        .collect();
    
    loop {
        steps += 1;
        let instruction = instructions[instruction_index];

        for current_node in current_nodes.iter_mut() {
            let node_to_find = match instruction {
                'L' => nodes[current_node.left_index.unwrap()].clone(),
                'R' => nodes[current_node.right_index.unwrap()].clone(),
                _ => panic!("invalid instruction {}", instruction),
            };
            *current_node = node_to_find;
        }

        if current_nodes
            .iter()
            .filter(|it| it.value.ends_with("Z"))
            .count()
            == current_nodes.len()
        {
            println!("found!");
            println!("{:?}", current_nodes);
            break;
        }

        if instruction_index == instructions.len() - 1 {
            instruction_index = 0;
            continue;
        }

        instruction_index += 1;
    }

    steps
}
