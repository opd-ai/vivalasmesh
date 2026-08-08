BEGIN {
    # flags
}
{
    line = $0
    # Remove trailing newline? awk already splits.
    # Determine if we need to add blank line before heading
    # We'll just print line after transformation, and ensure blank lines by adding empty print when needed.
    # Simpler: we will just transform and print; we can handle blank lines by checking previous line.
    # We'll store previous line.
}
