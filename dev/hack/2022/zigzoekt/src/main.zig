const std = @import("std");

const SimpleSection = struct{
    off: u32,
    sz: u32,
};

fn readSimpleSection(reader: anytype) !SimpleSection {
    const off = try reader.readInt(u32, std.builtin.Endian.Big);
    const sz = try reader.readInt(u32, std.builtin.Endian.Big);
    return SimpleSection{
        .off = off,
        .sz = sz,
    };
}

const CompoundSection = struct{
    data: SimpleSection,
    index: SimpleSection,
};

const TOC = struct{
    fileContents: CompoundSection,
    fileNames: CompoundSection,
};

fn readTOC(file: std.fs.File) !TOC {
    try file.seekFromEnd(-8);
    const tocSection = try readSimpleSection(file.reader());

    // Go to TOC
    try file.seekTo(tocSection.off);
    var limitReader = std.io.limitedReader(file.reader(), tocSection.sz);
    var reader = limitReader.reader();

    const sectionCount = try reader.readInt(u32, std.builtin.Endian.Big);
    if (sectionCount != 0) {
        // We only support 0
        return error.EndOfStream;
    }

    var toc: TOC = undefined;
    var buffer: [1024]u8 = undefined;
    while(limitReader.bytes_left > 0) {
        var slen = try std.leb.readULEB128(u64, reader);

        // Section Tag
        var name = buffer[0..slen];
        try reader.readNoEof(name);

        // Section Kind (0 = simple section, 1 = compound section)
        const kind = try reader.readByte();

        const data = try readSimpleSection(reader);
        const index: SimpleSection = try switch(kind) {
            0 => undefined,
            1, 2 => readSimpleSection(reader),
            else => undefined, // TODO return error. Couldn't get zig to work for me here.
        };

        if (std.mem.eql(u8, name, "fileContents")) {
            toc.fileContents = .{
                .data = data,
                .index = index,
            };
        } else if (std.mem.eql(u8, name, "fileNames")) {
            toc.fileNames = .{
                .data = data,
                .index = index,
            };
        }
    }

    return toc;
}

pub fn main() anyerror!void {
    const file = try std.fs.cwd().openFile(
        "github.com%2Fkeegancsmith%2Fsqlf_v16.00000.zoekt",
        .{},
    );
    defer file.close();

    const toc = try readTOC(file);

    try file.seekTo(toc.fileContents.data.off);
    var contentReader = std.io.limitedReader(file.reader(), toc.fileContents.data.sz).reader();

    var needle = "func";
    var buffer: [1024]u8 = undefined;
    while (contentReader.readUntilDelimiterOrEof(&buffer, '\n') catch { return; }) |line| {
        if (std.mem.containsAtLeast(u8, line, 1, needle)) {
            std.log.info("HI: {s}", .{line});
        }
    }
}
