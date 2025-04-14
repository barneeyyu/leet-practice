#!/usr/bin/env python3
# New script to generate README files
import os
import requests
import time


def get_problem_difficulty(problem_number):
    """Get problem difficulty from LeetCode API"""
    try:
        # Add delay to avoid rate limiting
        time.sleep(0.5)
        url = "https://leetcode.com/api/problems/all/"
        response = requests.get(url)
        if response.status_code == 200:
            data = response.json()
            for problem in data["stat_status_pairs"]:
                question_id = problem["stat"]["frontend_question_id"]
                if str(question_id) == problem_number:
                    level = problem["difficulty"]["level"]
                    return {1: "Easy", 2: "Medium", 3: "Hard"}.get(level, "Unknown")  # noqa: E501
    except Exception as e:
        print(f"Error fetching difficulty for problem {problem_number}: {e}")
    return "Unknown"


def generate_readme(folder_name):
    readme_path = os.path.join(folder_name, "README.md")
    problems = []

    for entry in os.listdir(folder_name):
        full_path = os.path.join(folder_name, entry)
        if os.path.isdir(full_path) and not entry.startswith("."):
            parts = entry.split("_", 1)
            if len(parts) == 2:
                problem_number = parts[0].strip()
                problem_title = parts[1].replace("_", " ").title()
                problem_difficulty = get_problem_difficulty(problem_number)
                solution_path = os.path.join(entry, "solution.go")
                notes_path = os.path.join(entry, "README.md")
                row = (
                    f"| {problem_number} | {problem_title} | "
                    f"{problem_difficulty} | [解法]({solution_path}) | "
                    f"[筆記]({notes_path}) |"
                )
                problems.append(row)

    folder_title = folder_name.replace("_", " ").title()
    readme_content = f"## {folder_title} 題型整理\n\n"
    readme_content += "| 題號 | 題目名稱 | 難度 | 解法 | 筆記 |\n"
    readme_content += "|------|----------|------|------|------|\n"
    readme_content += "\n".join(problems) + "\n"

    with open(readme_path, "w", encoding="utf-8") as f:
        f.write(readme_content)


def update_main_readme(stats):
    """Update top-level README.md difficulty badge section."""
    stats_block = (
        "<!-- stats-start -->\n"
        f"![Easy](https://img.shields.io/badge/Easy-{stats['Easy']}-green)\n"
        f"![Medium](https://img.shields.io/badge/Medium-{stats['Medium']}-yellow)\n"  # noqa: E501
        f"![Hard](https://img.shields.io/badge/Hard-{stats['Hard']}-red)\n"
        "<!-- stats-end -->"
    )

    readme_path = "README.md"
    with open(readme_path, "r+", encoding="utf-8") as f:
        content = f.read()
        start = content.find("<!-- stats-start -->")
        end = content.find("<!-- stats-end -->") + len("<!-- stats-end -->")
        if start != -1 and end != -1:
            new_content = content[:start] + stats_block + content[end:]
            f.seek(0)
            f.write(new_content)
            f.truncate()


def main():
    difficulty_count = {"Easy": 0, "Medium": 0, "Hard": 0}
    folders = [
        "array",
        "binary_search",
        "dfs_bfs",
        "dynamic_programming",
        "greedy",
        "hashmap",
        "sliding_window",
        "stack_queue",
        "string",
        "two_pointers",
    ]

    for folder in folders:
        if os.path.isdir(folder):
            generate_readme(folder)

            for entry in os.listdir(folder):
                full_path = os.path.join(folder, entry)
                if os.path.isdir(full_path) and not entry.startswith("."):
                    parts = entry.split("_", 1)
                    if len(parts) == 2:
                        problem_number = parts[0].strip()
                        difficulty = get_problem_difficulty(problem_number)
                        if difficulty in difficulty_count:
                            difficulty_count[difficulty] += 1

    update_main_readme(difficulty_count)
    print("Updated main README with difficulty stats.")


if __name__ == "__main__":
    main()
