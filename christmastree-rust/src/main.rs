use std::thread;
use std::time::Duration;

fn main() {
    thread::spawn(|| {
        println!("Server started at http://localhost:3000");
        draw(4); 
    });

    loop {
        thread::sleep(Duration::from_secs(1));
    }
}
fn draw(max: usize) {
    if max < 1 {
        println!("The number of levels must be at least 1.");
        return;
    }

    let total_width = (max * 2) - 1;
    let center = total_width / 2;

    for i in 0..max {
        let stars = "*".repeat(i * 2 + 1);
        let row = format!("{:>width$}",stars, width=center + i + 1);
        println!("{}", row);
    }

    let vertical_line = "|".to_string();
    let padded_vertical_line = format!("{:>width$}", vertical_line, width = center + 1);
    println!("{}", padded_vertical_line);
}
