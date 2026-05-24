import json
import os

def parse_notables():
    source_path = "data/trade_stats.json"
    output_path = "data/notables.json"
    
    if not os.path.exists(source_path):
        print(f"Error: {source_path} not found. Please run the curl command first.")
        return

    with open(source_path, "r", encoding="utf-8") as f:
        data = json.load(f)

    try:
        if "result" not in data:
            print("Error: The downloaded JSON is corrupt or missing the 'result' key.")
            return

        # Use a dictionary where the key is the clean name to automatically crush duplicates
        unique_notables = {}
        
        for section in data["result"]:
            entries = section.get("entries", [])
            for entry in entries:
                trade_id = entry.get("id")     
                name = entry.get("text")       
                
                if not name or not trade_id:
                    continue
                
                if "passive_notable_" in trade_id or "Added Passive Skill is" in name:
                    # Strip GGG's default structural search syntax to get the raw name
                    clean_name = name.replace("1 Added Passive Skill is ", "").strip()
                    
                    # DEDUPLICATION: If we already stored this passive name, skip the duplicate entry!
                    if clean_name in unique_notables:
                        continue

                    # Store temporarily using the name as the unique key
                    unique_notables[clean_name] = trade_id

        # Format into our clean, sequential integer indexed list for Go
        notables_pool = []
        unique_id = 0
        
        # Sort names alphabetically so your IDs stay completely deterministic
        for sorted_name in sorted(unique_notables.keys()):
            notables_pool.append({
                "id": unique_id,
                "name": sorted_name,
                "trade_id": unique_notables[sorted_name]
            })
            unique_id += 1

        if not notables_pool:
            print("Error: Universal sweep completed but found zero matching passive tree nodes.")
            return

        with open(output_path, "w", encoding="utf-8") as f:
            json.dump(notables_pool, f, indent=2)

        print(f"Successfully generated {output_path}!")
        print(f"Total DEDUPLICATED nodes mapped for your matrix layout: {len(notables_pool)}")

    except (KeyError, TypeError, IndexError) as e:
        print(f"CRITICAL ERROR: Structure mismatch inside trade_stats.json.")
        print(f"Details: {e}")
        return

if __name__ == "__main__":
    parse_notables()